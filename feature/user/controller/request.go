package controller

import "ecommerce-project/feature/user/entities"

type FotoUpdate struct {
	Foto string `json:"foto" form:"foto"`
}

type NameUpdate struct {
	Name string `json:"name" form:"name"`
}

type UserNameUpdate struct {
	Username string `json:"username" form:"username"`
}

type PasswordUpdate struct {
	Password string `json:"password" form:"password"`
}

type EmailUpdate struct {
	Email string `json:"email" form:"email"`
}

func UserNameToCore(data UserNameUpdate) entities.CoreUser {
	return entities.CoreUser{
		Username: data.Username,
	}
}

func PasswordToCore(data PasswordUpdate) entities.CoreUser {
	return entities.CoreUser{
		Password: data.Password,
	}
}

func EmailToCore(data EmailUpdate) entities.CoreUser {
	return entities.CoreUser{
		Email: data.Email,
	}
}

func NameToCore(data NameUpdate) entities.CoreUser {
	return entities.CoreUser{
		Name: data.Name,
	}
}
