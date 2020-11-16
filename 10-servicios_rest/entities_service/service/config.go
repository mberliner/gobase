package service

import "github.com/mberliner/gobase/10-servicios_rest/entities_service/repository"

var (
	//UserS para manejo del negocio de Usuario
	UserS UserService
	//PersonaS para manejo del negocio de Persona
	PersonaS PersonaService
)

func init() {

	UserS = NewUserService(repository.UserRepo)
	PersonaS = NewPersonaService(repository.PersonaRepo)

}
