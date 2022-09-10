package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ProductID []Product
	UserID    User
	Order     Order
}
