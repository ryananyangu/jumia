package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/ryananyangu/gonativeweb/controllers"
)

var base_path = "/api/v1"

var Routes = map[string]map[string]fiber.Handler{
	base_path + "/product/get": {
		http.MethodGet: controllers.GetProductBySKU,
	},
	base_path + "/product/update": {
		http.MethodPost: controllers.UpdateStock,
	},
	base_path + "/product/sell": {
		http.MethodPost: controllers.SellProducts,
	},
	base_path + "/product/bulk/update": {
		http.MethodPost: controllers.BulkUpdateStock,
	},
	base_path + "/product/threshold/update": {
		http.MethodPut: controllers.UpdateAlertThreshold,
	},
	base_path + "/product/stock/update/batchfile": {
		http.MethodPost: controllers.UpdateStockFromFile,
	},
}
