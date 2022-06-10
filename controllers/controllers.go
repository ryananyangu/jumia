package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/ryananyangu/jumia/models"
	"github.com/ryananyangu/jumia/services"
)

func GetProductBySKU(ctx *fiber.Ctx) error {

	sku := ctx.Query("sku")
	country_code := ctx.Query("country")

	product := services.GetProductSKUCountry(sku, country_code)
	return ctx.Status(http.StatusOK).JSON(product)
}

func SellProducts(ctx *fiber.Ctx) error {

	product := &models.Product{}
	err := ctx.BodyParser(product)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}
	result := services.MakeOrder(product)
	return ctx.Status(http.StatusOK).JSON(result)
}

func UpdateStock(ctx *fiber.Ctx) error {
	product := models.Product{}
	err := ctx.BodyParser(&product)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}
	services.ProductCreate(&product)
	return ctx.Status(http.StatusOK).JSON(product)

}

func BulkUpdateStock(ctx *fiber.Ctx) error {

	bulk := []models.Product{}
	err := ctx.BodyParser(&bulk)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}
	services.BatchUpdateProduct(&bulk)
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintln("Request recieved successfull"),
	})

}

func UpdateAlertThreshold(ctx *fiber.Ctx) error {
	product := models.Product{}
	err := ctx.BodyParser(&product)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}
	services.ProductTholdUpdate(&product)
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Product ID %d updated successfull", product.ID),
	})
}

func UpdateStockFromFile(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("fileUpload")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	buffer, err := file.Open()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	defer buffer.Close()
	reader := csv.NewReader(buffer)
	reader.Comma = ';'
	reader.LazyQuotes = true
	csvLines, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	services.BatchUpdateProductFile(csvLines)
	// for _, line := range csvLines {
	// 	utils.Info.Printf("%v", line)
	// 	break
	// }
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintln("Request recieved successfull"),
	})
}
