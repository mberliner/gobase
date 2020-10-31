package business

import (
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/repository"
)

var UserB UserBusiness
var PersonaB PersonaBusiness

func init() {

	UserB = NewUserBusiness(repository.UserRepo)
	PersonaB = NewPersonaBusiness(repository.PersonaRepo)

}
