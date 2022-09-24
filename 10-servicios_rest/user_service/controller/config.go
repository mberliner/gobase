package controller

import (
	"github.com/mberliner/gobase/10-servicios_rest/user_service/service"
)

var (
	//UserC controller de Usuario
	UserC UserController
)

func init() {

	UserC = NewUserController(service.UserS)

}
