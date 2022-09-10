package entities

import "time"

type CoreUser struct {
	ID         uint
	Name       string
	Email      string
	Password   string
	Address    string
	Foto       string
	Products   []CoreProduct
	Created_at time.Time
	Update_at  time.Time
}

type CoreProduct struct {
	ID          uint
	Name        string
	Price       int
	Description string
	Foto        string
	UserID      CoreUser
}
