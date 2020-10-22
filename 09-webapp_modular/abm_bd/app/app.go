package app

import (
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/controller"
	"log"
	"net/http"
)

func initServer() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func StartApp() {
	mapUrls()
	initServer()
}

func mapUrls() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/seccion", controller.Seccion)
	http.HandleFunc("/usuario", controller.AltaUser)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/persona/crear", controller.CrearPersona)
	http.HandleFunc("/persona", controller.AbmPersona)
	http.HandleFunc("/persona/borrar", controller.BorrarPersona)
	http.HandleFunc("/persona/editar", controller.ActualizarPersona)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}
