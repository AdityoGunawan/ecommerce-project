package controller

import "ecommerce-project/feature/product/entities"

type ResponseProduct struct {
	ID       uint
	Name     string
	Price    int
	Foto     string
	Quantity uint
}

func CoreToResponse(data entities.CoreProduct) ResponseProduct {
	var response ResponseProduct
	response.ID = data.ID
	response.Name = data.Name
	response.Price = data.Price
	response.Foto = data.Foto
	response.Quantity = data.Quantity

	return response
}

func CoreToResponseList(data []entities.CoreProduct) []ResponseProduct {
	var list []ResponseProduct
	for _, v := range data {
		temp := CoreToResponse(v)
		list = append(list, temp)
	}

	return list
}
