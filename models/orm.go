package models

import (
	auth "ecommerce-project/feature/auth/entities"
	product "ecommerce-project/feature/product/entities"
	"ecommerce-project/feature/user/entities"
)

func CoreToModel(core product.CoreProduct) Product {
	return Product{
		Name:        core.Name,
		Price:       core.Price,
		Description: core.Description,
		Foto:        core.Foto,
		UserID:      core.UserID,
		CategoryID:  core.CategoryID,
		Quantity:    core.Quantity,
	}
}

func RegCoreToModel(core auth.CoreRegister) User {
	return User{
		Username: core.Username,
		Email:    core.Email,
		Password: core.Password,
	}
}

func ModelToCore(model User) entities.CoreUser {
	return entities.CoreUser{
		Name:     model.Name,
		Email:    model.Email,
		Foto:     model.Foto,
		Username: model.Username,
	}
}

func ProductToCore(data Product) product.CoreProduct {
	return product.CoreProduct{
		ID:         data.ID,
		Name:       data.Name,
		Price:      data.Price,
		Foto:       data.Foto,
		UserID:     data.UserID,
		CategoryID: data.CategoryID,
		Quantity:   data.Quantity,
	}
}

func ProductToCoreList(data []Product) []product.CoreProduct {
	var corelist []product.CoreProduct
	for _, v := range data {
		temp := ProductToCore(v)
		corelist = append(corelist, temp)
	}

	return corelist
}

func CoreUserToModel(data entities.CoreUser) User {
	return User{
		Name:     data.Name,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Foto:     data.Foto,
	}
}
