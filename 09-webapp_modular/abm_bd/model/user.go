package model

import (
	"database/sql"
)

type User struct {
	ID       int
	Usuario  string
	Nombre   string
	Apellido string
	Edad     sql.NullInt64
	Password string
	Error    error
	Mensaje  string
}
