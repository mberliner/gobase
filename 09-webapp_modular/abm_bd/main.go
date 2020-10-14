package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/seccion", seccion)
	http.HandleFunc("/altausuario", altaUser)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/persona/crear", crearPersona)
	http.HandleFunc("/abmpersona", abmPersona)
	http.HandleFunc("/persona/borrar", borrarPersona)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
