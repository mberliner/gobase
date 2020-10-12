package main

import (
	"fmt"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/business"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/model"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/repository"
	"net/http"
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
		u = business.BuscaPorUsuario(usu)
	}
	return u
}

func estaLogueado(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	usu := dbSessions[c.Value]

	sU, err := repository.UR.BuscaPorUsuario(usu)
	fmt.Println("user:------------------estalogueado2-------------------------", sU, err, "len", len(sU))
	if err != nil || len(sU) == 0 {
		return false
	}
	return true
}
