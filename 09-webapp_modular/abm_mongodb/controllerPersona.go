package main

import (
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/business"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/model"
	_ "html/template"
	"log"
	"net/http"
	"strconv"
)

func abmPersona(res http.ResponseWriter, req *http.Request) {
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	mP := business.BuscaTodo()

	if err := tpl.ExecuteTemplate(res, "abmPersona.gohtml", mP); err != nil {
		log.Println("Error en abmPersona:", err)
	}
}

func crearPersona(res http.ResponseWriter, req *http.Request) {
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var mP model.Personas

	if req.Method == http.MethodPost {

		nom := req.FormValue("nombre")
		ape := req.FormValue("apellido")
		fecha := req.FormValue("fechaNacimiento")

		mP1 := business.CreaPersona(nom, ape, fecha)
		//TODO Revisar si eliminar
		if mP.Error != nil {
			if err := tpl.ExecuteTemplate(res, "crearPersona.gohtml", mP1); err != nil {
				log.Println("Error en crearPersona:", err)
			}
			return
		}

		mP = business.BuscaTodo()
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

func borrarPersona(res http.ResponseWriter, req *http.Request) {
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	idT := req.FormValue("id")
	id, _ := strconv.Atoi(idT)
	mP1 := business.BorraPersona(id)

	mP := business.BuscaTodo()
	if mP.Error == nil {
		mP.Error = mP1.Error
		mP.Mensaje = mP1.Mensaje
	}

	if err := tpl.ExecuteTemplate(res, "abmPersona.gohtml", mP); err != nil {
		log.Println("Error en abmPersona al borrar:", err)
	}
}

func actualizarPersona(res http.ResponseWriter, req *http.Request) {
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var mP model.Personas
	idT := req.FormValue("id")

	if req.Method == http.MethodPost {

		//id := req.FormValue("id")
		nom := req.FormValue("nombre")
		ape := req.FormValue("apellido")
		fecha := req.FormValue("fechaNacimiento")

		mP1 := business.ActualizaPersona(idT, nom, ape, fecha)
		if mP.Error != nil {
			if err := tpl.ExecuteTemplate(res, "editarPersona.gohtml", mP1); err != nil {
				log.Println("Error en editarPersona:", err)
			}
			return
		}

		mP = business.BuscaTodo()
		if mP.Error == nil {
			mP.Error = mP1.Error
			mP.Mensaje = mP1.Mensaje
		}

		if err := tpl.ExecuteTemplate(res, "abmPersona.gohtml", mP); err != nil {
			log.Println("Error en abmPersona:", err)
		}
		return

	}

	id, _ := strconv.Atoi(idT)
	mP = business.BuscaPersona(id)

	if err := tpl.ExecuteTemplate(res, "editarPersona.gohtml", mP); err != nil {
		log.Println("Error en editarPersona:", err)
	}
}
