package service

import "github.com/mberliner/gobase/10-servicios_rest/user_service/repository"

var (
	//UserS para manejo del negocio de Usuario
	UserS UserService
)

func init() {

	UserS = NewUserService(repository.UserRepo)

}
