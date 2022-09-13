package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Price       int
	Description string
	Foto        string
	UserID      uint
	CategoryID  uint
	Quantity    uint
	Carts       []Cart
	Orders      []Order
}

type Category struct {
	gorm.Model
	Category string
	Products []Product
}
