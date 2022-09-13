package entities

type ServiceInterface interface {
	AddProductI(CoreProduct) (int, error)
	GetAll() ([]CoreProduct, error)
	GetMyProduct(id int) ([]CoreProduct, error)
	Delete(userid, deleteid int) (string, error)
}

type RepositoryInterface interface {
	InsertI(CoreProduct) (int, error)
	SelectAll() ([]CoreProduct, error)
	SelectMyProduct(id int) ([]CoreProduct, error)
	Delete(userid, deleteid int) (string, error)
}
