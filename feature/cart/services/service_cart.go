package services

import (
	"ecommerce-project/config"
	"ecommerce-project/models"
)

func CreateCart(Cart *models.Cart) (interface{}, error) {

	if err := config.DB.Create(&Cart).Error; err != nil {
		return nil, err
	}

	return Cart.UserID, nil
}
