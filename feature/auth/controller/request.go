package controller

import "ecommerce-project/feature/auth/entities"

type RegisterReq struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginReq struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (request *RegisterReq) ReqToCore() entities.CoreRegister {
	core := entities.CoreRegister{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
	return core
}

func (request *LoginReq) ReqToCoreLogin() entities.CoreLogin {
	core := entities.CoreLogin{
		Email:    request.Email,
		Password: request.Password,
	}
	return core
}
