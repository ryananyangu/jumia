package services

import (
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/ryananyangu/gonativeweb/models"
	"github.com/ryananyangu/gonativeweb/utils"
)

var wg sync.WaitGroup

func ProductTholdUpdate(product *models.Product) *models.Product {
	// Update product to reflect the stock change
	utils.Db.Model(models.Product{}).Where("id = ?", product.ID).
		Updates(models.Product{AlertThold: product.AlertThold})
	CheckAlertThold(product)
	return product
}

func GetProductSKUCountry(sku, country string) *models.Product {
	product := &models.Product{}
	utils.Db.Where("sku = ? AND country = ?", sku, country).
		First(product)
	return product

}
func ProductCreate(product *models.Product) *models.Product {
	stockChange := int(product.Stock)
	tx := utils.Db.Begin()
	tx.Where("sku = ? AND country = ?", product.SKU, product.Country).
		First(product)

	if product.ID == 0 {
		// Create product
		if product.AlertThold <= 0 {
			product.AlertThold = utils.DEFAULT_ALERT_THOLD
		}
		if err := tx.Create(product).Error; err != nil {
			tx.Rollback()
			utils.Error.Println(err.Error())
			return product
		}

		tx.Commit()

		utils.Info.Printf("Stocked Product : %s, Country : %s, Sku : %s, Stock : %d ",
			product.Name,
			product.Country,
			product.SKU,
			stockChange,
		)
		return product
	}

	// Update product to reflect the stock change
	if err := tx.Model(models.Product{}).Where("id = ?", product.ID).
		Updates(models.Product{Stock: product.Stock + stockChange}).Error; err != nil {
		tx.Rollback()
		utils.Error.Println(err.Error())
		return product
	}
	product.Stock += stockChange

	tx.Commit()
	utils.Info.Printf("Stocked Product : %s, Country : %s, Sku : %s, Stock : %d ",
		product.Name,
		product.Country,
		product.SKU,
		stockChange,
	)
	return product
}

func BatchUpdateProduct(products *[]models.Product) {

	dbconn_guard := make(chan int, utils.MAX_DB_CONNECTIONS)

	for _, product := range *products {
		// Async Batch Processing of the products
		wg.Add(1)
		dbconn_guard <- 1 // will block if there is MAX_DB_CONNECTIONS ints in dbconn_guard
		go func(product models.Product) {
			SwitchSellBuy(&product)
			wg.Done()
			<-dbconn_guard

		}(product)

	}

	// wg.Wait() // Commented out to run some jobs in the background
}

func BatchUpdateProductFile(rawInput [][]string) {

	dbconn_guard := make(chan int, utils.MAX_DB_CONNECTIONS)
	for index, inputLine := range rawInput {

		// Avoid inserting the headers into the DB
		if index == 0 {
			continue
		}

		// cleaning stray double quotes
		res := strings.ReplaceAll(inputLine[0], "\",\"", "|")

		// Spliting the file based on commas
		resArry := strings.Split(res, "|")

		// Validate columns in csv
		if len(resArry) != 4 {
			utils.Error.Printf("%v | Invalid file line", inputLine)
			continue
		}

		// validate stock amount
		stk, err := strconv.Atoi(resArry[3])
		if err != nil {
			utils.Error.Printf("%v | %s | Stock Amount value invalid", inputLine, err.Error())
			continue
		}

		product := &models.Product{
			Country: resArry[0],
			SKU:     resArry[1],
			Name:    resArry[2],
			Stock:   stk,
		}
		wg.Add(1)
		dbconn_guard <- 1 // will block if there is MAX_DB_CONNECTIONS ints in dbconn_guard
		go func(product models.Product) {
			SwitchSellBuy(&product)
			wg.Done()
			<-dbconn_guard
		}(*product)

	}

	// wg.Wait() Commented out to run fully on the back ground

}

func MakeOrder(product *models.Product) *models.Order {

	stock_Amnt := int(math.Abs(float64(product.Stock)))

	utils.Db.First(product, "sku = ? AND country = ? AND stock >= ?",
		product.SKU,
		product.Country,
		stock_Amnt,
	)

	order := models.Order{
		ProductID: product.ID,
		Amount:    uint(stock_Amnt),
	}

	if product.ID <= 0 {
		utils.Error.Printf("%v | Product Un-Available or stock not enough for order", order)
		return &order

	}

	// Initialize order
	utils.Db.Create(&order)

	// Update product to reflect the stock change
	utils.Db.Model(models.Product{}).Where("id = ?", product.ID).
		Updates(models.Product{Stock: product.Stock - stock_Amnt})

	// Check to send alert
	CheckAlertThold(product)

	utils.Info.Printf("Sold Product : %s, Country : %s, Sku : %s, Stock : %d ",
		product.Name,
		product.Country,
		product.SKU,
		stock_Amnt,
	)

	return &order

}

// Log threshold alert to stdout
func CheckAlertThold(product *models.Product) {
	if int(product.AlertThold) >= product.Stock {
		utils.Warning.Printf("%s SKU = %s COUNTRY = %s alert threshold of : %d Reached/Below current value %d ",
			product.Name,
			product.SKU,
			product.Country,
			product.AlertThold,
			product.Stock,
		)
	}
}

// Switch used in batch updates to switch between order and stock
func SwitchSellBuy(product *models.Product) {
	if int(product.Stock) < 0 {
		MakeOrder(product)
	} else {
		ProductCreate(product)
	}

}
