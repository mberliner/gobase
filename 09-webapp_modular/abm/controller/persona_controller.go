package controller

import (
	"log"
	"net/http"

	"github.com/mberliner/gobase/09-webapp_modular/abm/model"
	"github.com/mberliner/gobase/09-webapp_modular/abm/service"
)

func AbmPersona(res http.ResponseWriter, req *http.Request) {
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	mP := service.PersonaS.BuscaTodo()

	if err := tpl.ExecuteTemplate(res, "abmPersona.gohtml", mP); err != nil {
		log.Println("Error en abmPersona:", err)
	}
}

func CrearPersona(res http.ResponseWriter, req *http.Request) {
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var mP model.Personas

	if req.Method == http.MethodPost {

		nom := req.FormValue("nombre")
		ape := req.FormValue("apellido")
		fecha := req.FormValue("fechaNacimiento")

		mP1 := service.PersonaS.CreaPersona(nom, ape, fecha)
		if mP.Error != nil {
			if err := tpl.ExecuteTemplate(res, "crearPersona.gohtml", mP1); err != nil {
				log.Println("Error en crearPersona:", err)
			}
			return
		}

		mP = service.PersonaS.BuscaTodo()
		if mP.Error == nil {
			mP.Error = mP1.Error
			mP.Mensaje = mP1.Mensaje
		}

		if err := tpl.ExecuteTemplate(res, "abmPersona.gohtml", mP); err != nil {
			log.Println("Error en abmPersona:", err)
		}
		return

	}

	if err := tpl.ExecuteTemplate(res, "crearPersona.gohtml", mP); err != nil {
		log.Println("Error en crearPersona:", err)
	}
}

func BorrarPersona(res http.ResponseWriter, req *http.Request) {
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	id := req.FormValue("id")
	mP1 := service.PersonaS.BorraPersona(id)

	mP := service.PersonaS.BuscaTodo()
	if mP.Error == nil {
		mP.Error = mP1.Error
		mP.Mensaje = mP1.Mensaje
	}

	if err := tpl.ExecuteTemplate(res, "abmPersona.gohtml", mP); err != nil {
		log.Println("Error en abmPersona al borrar:", err)
	}
}

func ActualizarPersona(res http.ResponseWriter, req *http.Request) {
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var mP model.Personas
	id := req.FormValue("id")

	if req.Method == http.MethodPost {

		nom := req.FormValue("nombre")
		ape := req.FormValue("apellido")
		fecha := req.FormValue("fechaNacimiento")

		mP1 := service.PersonaS.ActualizaPersona(id, nom, ape, fecha)
		if mP.Error != nil {
			if err := tpl.ExecuteTemplate(res, "editarPersona.gohtml", mP1); err != nil {
				log.Println("Error en editarPersona:", err)
			}
			return
		}

		mP = service.PersonaS.BuscaTodo()
		if mP.Error == nil {
			mP.Error = mP1.Error
			mP.Mensaje = mP1.Mensaje
		}

		if err := tpl.ExecuteTemplate(res, "abmPersona.gohtml", mP); err != nil {
			log.Println("Error en abmPersona:", err)
		}
		return

	}

	mP = service.PersonaS.BuscaPersona(id)

	if err := tpl.ExecuteTemplate(res, "editarPersona.gohtml", mP); err != nil {
		log.Println("Error en editarPersona:", err)
	}
}
