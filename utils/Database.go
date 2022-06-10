package utils

import (
	"fmt"
	"os"

	"github.com/ryananyangu/gonativeweb/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// : db
// : 5432
// : postgres
// : postgres
// : postgres
var Db *gorm.DB

func init() {

	if os.Getenv("ENV") == PROD_ENV {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Monrovia",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		// dsn := ""
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic(err)
		}
		Db = db
	} else {
		db, err := gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic(err)
		}
		Db = db

	}

	// Migrate the schema
	Db.AutoMigrate(models.Order{}, models.Product{})
}
