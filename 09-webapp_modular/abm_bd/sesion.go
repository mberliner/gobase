package main

import (
	"fmt"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/repository"
	"net/http"
)

var dbSessions = map[string]string{}

func getUser(res http.ResponseWriter, req *http.Request) repository.User {
	var u repository.User

	c, err := req.Cookie("session")
	if err != nil {
		return u
	}
	// Si existe lo tomo de la sesion
	if usu, ok := dbSessions[c.Value]; ok {
		sU, err := repository.UR.BuscaPorUsuario(usu)
		fmt.Println("user:------------------getUser3-------------------------", sU, err)
		if err != nil || len(sU) == 0 {
			//TODO log y revisar
			return u
		}
		u = sU[0]
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
	//	_, ok := dbUsers[un]
	return true
}
