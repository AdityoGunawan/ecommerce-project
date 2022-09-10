package entities

type CoreCart struct {
	ID        uint
	ProductID CoreProduct
	UserID    CoreUser
}

type CoreProduct struct {
	ID          uint
	Name        string
	Price       int
	Description string
	Foto        string
}
type CoreUser struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Address  string
}
