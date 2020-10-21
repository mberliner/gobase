package model

type User struct {
	ID       int
	Usuario  string
	Nombre   string
	Apellido string
	Edad     string
	Password string
	Error    error
	Mensaje  string
}
