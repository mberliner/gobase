package service

import "github.com/mberliner/gobase/09-webapp_modular/abm_bd/repository/mysql"

var (
	//UserS para manejo del negocio de Usuario
	UserS UserService
	//PersonaS para manejo del negocio de Persona
	PersonaS PersonaService
)

func init() {

	UserS = NewUserService(mysql.UserRepo)
	PersonaS = NewPersonaService(mysql.PersonaRepo)

}
