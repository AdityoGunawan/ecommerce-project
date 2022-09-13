package service

import (
	"ecommerce-project/feature/auth/entities"
)

type Service struct {
	do entities.RepositoryInterface
}

func New(data entities.RepositoryInterface) entities.ServiceInterface {
	return &Service{
		do: data,
	}
}

func (Service *Service) FirstRegister(core entities.CoreRegister) (int, error) {
	hash, _ := HashPassword(core.Password)
	core.Password = hash
	row, err := Service.do.InsertInto(core)
	return row, err
}

func (Service *Service) Login(email, password string) (string, error) {
	token, err := Service.do.Generate(email, password)
	return token, err
}
