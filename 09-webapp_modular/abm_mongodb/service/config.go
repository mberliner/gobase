package service

import (
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/repository/mongoDB"
)

var (
	//UserS para manejo del negocio de Usuario
	UserS UserService
	//PersonaS para manejo del negocio de Persona
	PersonaS PersonaService
)

func init() {

	UserS = NewUserService(mongoDB.UserRepo)
	PersonaS = NewPersonaService(mongoDB.PersonaRepo)

}
