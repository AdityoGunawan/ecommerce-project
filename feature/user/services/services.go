package services

import (
	"ecommerce-project/feature/user/entities"
)

type Services struct {
	do entities.RepositoryInterface
}

func New(data entities.RepositoryInterface) entities.ServiceInterface {
	return &Services{
		do: data,
	}
}

func (service *Services) MyProfile(id int) (entities.CoreUser, error) {
	core, err := service.do.Select(id)
	return core, err
}

func (service *Services) Update(data entities.CoreUser, id int) (string, error) {
	msg, err := service.do.Update(data, id)
	return msg, err
}

func (service *Services) Delete(userid int) (string, error) {
	msg, err := service.do.Delete(userid)
	return msg, err
}
