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

func altaPersona(res http.ResponseWriter, req *http.Request) {
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var p model.Personas

	if req.Method == http.MethodPost {

		nom := req.FormValue("nombre")
		ape := req.FormValue("apellido")
		fecha := req.FormValue("fechaNacimiento")

		p = business.CreaPersona(nom, ape, fecha)
		if p.Error != nil {
			if err := tpl.ExecuteTemplate(res, "altaPersona.gohtml", p); err != nil {
				log.Println("Error en altaUser:", err)
			}
			return
		}

	}

	if err := tpl.ExecuteTemplate(res, "altapersona.gohtml", p); err != nil {
		log.Println("Error en altapersona:", err)
	}
}

func borraPersona(res http.ResponseWriter, req *http.Request) {
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
