package model

type Persona struct {
	ID              string
	Nombre          string
	Apellido        string
	FechaNacimiento string
}

type Personas struct {
	PersonasM []Persona
	Error     error
	Mensaje   string
}
