package entities

type ServiceInterface interface {
	FirstRegister(CoreRegister) (int, error)
	Login(email, password string) (string, error)
}

type RepositoryInterface interface {
	InsertInto(CoreRegister) (int, error)
	Generate(email, password string) (string, error)
}
