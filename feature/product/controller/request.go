package controller

import "ecommerce-project/feature/product/entities"

type Request struct {
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Foto        string `json:"foto" form:"foto"`
	Description string `json:"description" form:"description"`
	Category    uint   `json:"category" form:"category"`
	Quantity    uint   `json:"quantity" form:"quantity"`
}

func (Req *Request) ReqToCore(foto string, userid uint) entities.CoreProduct {
	core := entities.CoreProduct{
		Name:        Req.Name,
		Price:       Req.Price,
		Foto:        foto,
		Description: Req.Description,
		CategoryID:  Req.Category,
		Quantity:    Req.Quantity,
		UserID:      userid,
	}

	return core
}
