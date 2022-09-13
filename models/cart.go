package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ProductID uint
	UserID    uint
	Product   Product
	Quantity  int
	Order     Order
}
