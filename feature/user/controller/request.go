package controller

import "ecommerce-project/feature/user/entities"

type Request struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Foto     string `json:"foto" form:"foto"`
}

func (ex *Request) RequestToCore() entities.CoreUser {
	var core entities.CoreUser
	core.Name = ex.Name
	core.Username = ex.Username
	core.Email = ex.Email
	core.Password = ex.Password
	core.Foto = ex.Foto

	return core
}
