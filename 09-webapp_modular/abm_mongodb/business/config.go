package business

import (
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/repository"
)

var UserB UserBusiness

func init() {

	UserB = NewUserBusiness(repository.UserRepo)

}
