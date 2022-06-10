package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Product   Product
	ProductID uint
	Amount    uint
}
