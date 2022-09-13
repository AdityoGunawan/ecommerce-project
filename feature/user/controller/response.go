package controller

import "ecommerce-project/feature/user/entities"

type Response struct {
	Name     string
	Email    string
	Foto     string
	Username string
}

func CoreToResponse(core entities.CoreUser) Response {
	return Response{
		Name:     core.Name,
		Email:    core.Email,
		Foto:     core.Foto,
		Username: core.Username,
	}
}
