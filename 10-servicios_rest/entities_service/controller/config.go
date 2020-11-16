package controller

import (
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/service"
)

var (
	//UserC controller de Usuario
	UserC UserController
	//PersonaC controller de Persona
	PersonaC PersonaController
)

func init() {

	UserC = NewUserController(service.UserS)
	PersonaC = NewPersonaController(service.PersonaS)

}
