package controller

import (
	"log"
	"net/http"

	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/business"
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/model"
)

func AbmPersona(res http.ResponseWriter, req *http.Request) {
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	mP := business.PersonaB.BuscaTodo()

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

		mP1 := business.PersonaB.CreaPersona(nom, ape, fecha)
		if mP.Error != nil {
			if err := tpl.ExecuteTemplate(res, "crearPersona.gohtml", mP1); err != nil {
				log.Println("Error en crearPersona:", err)
			}
			return
		}

		mP = business.PersonaB.BuscaTodo()
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
	mP1 := business.PersonaB.BorraPersona(id)

	mP := business.PersonaB.BuscaTodo()
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

		mP1 := business.PersonaB.ActualizaPersona(id, nom, ape, fecha)
		if mP.Error != nil {
			if err := tpl.ExecuteTemplate(res, "editarPersona.gohtml", mP1); err != nil {
				log.Println("Error en editarPersona:", err)
			}
			return
		}

		mP = business.PersonaB.BuscaTodo()
		if mP.Error == nil {
			mP.Error = mP1.Error
			mP.Mensaje = mP1.Mensaje
		}

		if err := tpl.ExecuteTemplate(res, "abmPersona.gohtml", mP); err != nil {
			log.Println("Error en abmPersona:", err)
		}
		return

	}

	mP = business.PersonaB.BuscaPersona(id)

	if err := tpl.ExecuteTemplate(res, "editarPersona.gohtml", mP); err != nil {
		log.Println("Error en editarPersona:", err)
	}
}
