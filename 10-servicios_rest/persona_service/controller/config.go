package controller

import (
	"github.com/mberliner/gobase/10-servicios_rest/persona_service/service"
)

var (
	//PersonaC controller de Persona
	PersonaC PersonaController
)

func init() {

	PersonaC = NewPersonaController(service.PersonaS)

}
