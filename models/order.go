package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID        uint
	ProductID     uint
	CartID        uint
	SendAddressID uint
	PaymentMethod string
}

type SendAddress struct {
	gorm.Model
	SendAddress string
	OrderID     []Order
}
