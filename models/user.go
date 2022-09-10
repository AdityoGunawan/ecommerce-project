package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Address  string
	Foto     string
	Products []Product
}

// `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
