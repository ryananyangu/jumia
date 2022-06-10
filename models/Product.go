package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string `json:"name"`
	SKU        string `gorm:"index:idx_skucountry,priority:1" json:"sku"`
	Stock      int    `json:"stock"`
	Country    string `gorm:"index:idx_skucountry,priority:2" json:"country"`
	AlertThold uint   `json:"alertThold"`
}
