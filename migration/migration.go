package migration

import (
	model "ecommerce-project/models"

	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Cart{})
	db.AutoMigrate(&model.Order{})
	db.AutoMigrate(&model.AddressRequest{})
	db.AutoMigrate(&model.ListOrder{})

}
