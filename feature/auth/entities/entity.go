package entities

type CoreRegister struct {
	Username string
	Email    string
	Password string
}

type CoreLogin struct {
	Email    string
	Password string
}
