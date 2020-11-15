package service

import "github.com/mberliner/gobase/10-servicios_rest/entities_service/repository"

var (
	//UserS para manejo del negocio de Usuario
	UserS UserService
	//PersonaB para manejo del negocio de Persona
	PersonaB PersonaBusiness
)

func init() {

	UserS = NewUserService(repository.UserRepo)
	PersonaB = NewPersonaBusiness(repository.PersonaRepo)

}
