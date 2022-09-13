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

func (service *Services) UpdateImage(alamat string, id int) (string, error) {
	msg, err := service.do.UpdateImage(alamat, id)
	return msg, err
}

func (service *Services) UpdateUserName(newname entities.CoreUser, id int) (string, error) {
	msg, err := service.do.UpdateUserName(newname, id)
	return msg, err
}

func (service *Services) UpdatePassword(NewPassword entities.CoreUser, id int) (string, error) {
	msg, err := service.do.UpdatePassword(NewPassword, id)
	return msg, err
}

func (service *Services) UpdateEmail(NewPassword entities.CoreUser, id int) (string, error) {
	msg, err := service.do.UpdateEmail(NewPassword, id)
	return msg, err
}

func (service *Services) UpdateName(NewPassword entities.CoreUser, id int) (string, error) {
	msg, err := service.do.UpdateName(NewPassword, id)
	return msg, err
}

func (service *Services) Delete(userid int) (string, error) {
	msg, err := service.do.Delete(userid)
	return msg, err
}
