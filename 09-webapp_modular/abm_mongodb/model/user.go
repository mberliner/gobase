package model

type User struct {
	ID       string
	Usuario  string
	Nombre   string
	Apellido string
	Edad     string
	Password string
	Error    error
	Mensaje  string
}
