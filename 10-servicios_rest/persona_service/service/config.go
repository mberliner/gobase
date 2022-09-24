package service

import "github.com/mberliner/gobase/10-servicios_rest/persona_service/repository"

var (
	//PersonaS para manejo del negocio de Persona
	PersonaS PersonaService
)

func init() {

	PersonaS = NewPersonaService(repository.PersonaRepo)

}
