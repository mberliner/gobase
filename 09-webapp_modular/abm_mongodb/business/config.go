package business

import (
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/repository/mongoDB"
)

var (
	//UserB para manejo del negocio de Usuario
	UserB UserBusiness
	//PersonaB para manejo del negocio de Persona
	PersonaB PersonaBusiness
)

func init() {

	UserB = NewUserBusiness(mongoDB.UserRepo)
	PersonaB = NewPersonaBusiness(mongoDB.PersonaRepo)

}
