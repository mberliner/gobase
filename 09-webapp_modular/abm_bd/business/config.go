package business

import "github.com/mberliner/gobase/09-webapp_modular/abm_bd/repository"

var (
	//UserB para manejo del negocio de Usuario
	UserB UserBusiness
	//PersonaB para manejo del negocio de Persona
	PersonaB PersonaBusiness
)

func init() {

	UserB = NewUserBusiness(repository.UserRepo)
	PersonaB = NewPersonaBusiness(repository.PersonaRepo)

}
