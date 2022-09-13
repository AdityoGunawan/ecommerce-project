package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID         uint
	ProductID      uint
	TotalQty       int
	TotalPrice     int
	AddressRequest uint
	PaymentMethod  string
	ListOrder      []ListOrder
}
type ListOrder struct {
	gorm.Model
	OrderID uint
	CartID  uint
}
type AddressRequest struct {
	gorm.Model
	Street string
	City   string
	State  string
	Zip    int
	Order  Order
}
type OrderRequest struct {
	Order   Order
	Address AddressRequest
}

// type SendAddress struct {
// 	gorm.Model
// 	SendAddress string
// 	OrderID     []Order
// }
