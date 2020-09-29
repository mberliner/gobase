package main

import (
	"net/http"
)

func getUser(req *http.Request) user {
	var u user

	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	// Si existe lo tomo de la sesion
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func estaLogueado(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
