package entities

type ServiceInterface interface {
	MyProfile(id int) (CoreUser, error)
	UpdateImage(image string, id int) (string, error)
	UpdateName(name CoreUser, id int) (string, error)
	UpdateUserName(name CoreUser, id int) (string, error)
	UpdatePassword(paswword CoreUser, id int) (string, error)
	UpdateEmail(email CoreUser, id int) (string, error)
	Delete(id int) (string, error)
}

type RepositoryInterface interface {
	Select(id int) (CoreUser, error)
	UpdateImage(image string, id int) (string, error)
	UpdateName(name CoreUser, id int) (string, error)
	UpdateUserName(name CoreUser, id int) (string, error)
	UpdatePassword(paswword CoreUser, id int) (string, error)
	UpdateEmail(email CoreUser, id int) (string, error)
	Delete(id int) (string, error)
}
