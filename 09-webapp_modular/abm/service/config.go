package service

import mongodb "github.com/mberliner/gobase/09-webapp_modular/abm/repository/mongoDB"

var (
	//UserS para manejo del negocio de Usuario
	UserS UserService
	//PersonaS para manejo del negocio de Persona
	PersonaS PersonaService
)

func init() {

	UserS = NewUserService(mongodb.UserRepo)
	PersonaS = NewPersonaService(mongodb.PersonaRepo)

}
