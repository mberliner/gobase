package model

type Persona struct {
	ID              int
	Nombre          string
	Apellido        string
	FechaNacimiento string
}

type Personas struct {
	PersonasM []Persona
	Error     error
	Mensaje   string
}
