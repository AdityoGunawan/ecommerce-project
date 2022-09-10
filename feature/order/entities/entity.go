package entities

import "time"

type CoreOrder struct {
	ID            uint
	UserID        CoreUser
	ProductID     CoreProduct
	CartID        CoreCart
	PaymentMethod string
	Created_at    time.Time
}

type CoreProduct struct {
	ID          uint
	Name        string
	Price       int
	Description string
	Foto        string
}

type CoreCart struct {
	ID        uint
	ProductID CoreProduct
}

type CoreUser struct {
	ID       uint
	Name     string
	Email    string
	Password string
}

type SendAddress struct {
	ID          uint
	SendAddress string
}
