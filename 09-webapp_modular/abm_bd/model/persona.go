package model

import (
	"database/sql"
)

type Persona struct {
	ID              int
	Nombre          string
	Apellido        string
	FechaNacimiento sql.NullTime
}

type Personas struct {
	PersonasM []Persona
	Error     error
	Mensaje   string
}
