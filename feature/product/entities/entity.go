package entities

import "time"

type CoreProduct struct {
	ID          uint
	Name        string
	Price       int
	Description string
	Foto        string
	UserID      uint
	CategoryID  uint
	Quantity    uint
	Created_at  time.Time
	Update_at   time.Time
}

type CoreCategory struct {
	ID         uint
	Category   string
	Created_at time.Time
	Update_at  time.Time
}

type CoreUser struct {
	ID         uint
	Name       string
	Email      string
	Password   string
	Address    string
	Foto       string
	Created_at time.Time
	Update_at  time.Time
}
