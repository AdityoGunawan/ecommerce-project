package services

import "ecommerce-project/feature/product/entities"

type Service struct {
	do entities.RepositoryInterface
}

func New(data entities.RepositoryInterface) entities.ServiceInterface {
	return &Service{
		do: data,
	}
}

func (Service *Service) AddProductI(core entities.CoreProduct) (int, error) {
	row, err := Service.do.InsertI(core)
	if err != nil {
		return 0, err
	}
	return row, err
}

func (service *Service) GetAll(page int) ([]entities.CoreProduct, error) {
	all, err := service.do.SelectAll(page)
	return all, err
}

func (service *Service) Delete(userid, deleteid int) (string, error) {
	msg, err := service.do.Delete(userid, deleteid)
	return msg, err
}

func (service *Service) GetMyProduct(userid int) ([]entities.CoreProduct, error) {
	msg, err := service.do.SelectMyProduct(userid)
	return msg, err
}
