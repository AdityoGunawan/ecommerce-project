package entities

type ServiceInterface interface {
	MyProfile(id int) (CoreUser, error)
	Update(data CoreUser, id int) (string, error)
	Delete(id int) (string, error)
}

type RepositoryInterface interface {
	Select(id int) (CoreUser, error)
	Update(data CoreUser, id int) (string, error)
	Delete(id int) (string, error)
}
