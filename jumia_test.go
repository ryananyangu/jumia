package main

import (
	"net/http"
	"strings"
	"testing"

	// "github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
)

func TestProductCreate(t *testing.T) {
	payload := `{
		"name":"Samsung Phone",
		"sku" : "UYUT-879847564793-PO",
		"stock":2,
		"alertThold":10,
		"country":"ke"
	}`

	req, _ := http.NewRequest("POST", "/api/v1/product/update", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")

	res, _ := SetupRouter().Test(req, -1)
	assert.Equal(t, 200, res.StatusCode)

}

func TestProductGet(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/v1/product/get?sku=UYUT-879847564793-PO&country=ke", nil)
	req.Header.Add("Content-Type", "application/json")

	res, _ := SetupRouter().Test(req, -1)
	assert.Equal(t, 200, res.StatusCode)

}

func TestProductOrder(t *testing.T) {
	payload := `{
		"sku": "UYUT-879847564793-PO",
		"stock": -1,
		"country": "ke"
	}`

	req, _ := http.NewRequest("POST", "/api/v1/product/sell", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")

	res, _ := SetupRouter().Test(req, -1)
	assert.Equal(t, 200, res.StatusCode)

}

func TestProductBulk(t *testing.T) {
	payload := `[
		{
			"name": "Samsung Phone",
			"sku": "UYUT-879847564793-PO",
			"stock": 14,
			"country": "ke",
			"alertThold": 10
		},
		{
			"name": "Samsung TAB",
			"sku": "UYUT-879847564793-TAB",
			"stock": 200,
			"country": "ke",
			"alertThold": 10
		},
		{
			"name": "Samsung TV",
			"sku": "UYUT-879847564793-TV",
			"stock": 5,
			"country": "ke",
			"alertThold": 10
		},
		{
			"name": "Samsung Phone",
			"sku": "UYUT-879847564793-PO",
			"stock": -6,
			"country": "ke",
			"alertThold": 10
		}
	]`

	req, _ := http.NewRequest("POST", "/api/v1/product/bulk/update", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")

	res, _ := SetupRouter().Test(req, -1)
	assert.Equal(t, 200, res.StatusCode)

}

func TestProductThreshholdUpdate(t *testing.T) {
	payload := `{
		"ID" : 1,
		"alertThold":20
	}`

	req, _ := http.NewRequest("POST", "/api/v1/product/threshold/update", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")

	res, _ := SetupRouter().Test(req, -1)
	assert.Equal(t, 200, res.StatusCode)

}
