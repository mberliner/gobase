package controller

import (
	"net/http"

	"github.com/mberliner/gobase/09-webapp_modular/abm/model"
	"github.com/mberliner/gobase/09-webapp_modular/abm/service"
)

const sessionCookie string = "session"

var dbSessions = map[string]string{}

func getUser(res http.ResponseWriter, req *http.Request) model.User {
	var u model.User

	c, err := req.Cookie("session")
	if err != nil {
		return u
	}
	// Si existe lo tomo de la sesion
	if usu, ok := dbSessions[c.Value]; ok {
		u = service.UserS.BuscaPorUsuario(usu)
	}
	return u
}

func estaLogueado(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	usu := dbSessions[c.Value]

	u := service.UserS.BuscaPorUsuario(usu)
	if u.Error != nil {
		return false
	}
	return true
}
