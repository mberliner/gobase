package main

import (
	"github.com/google/uuid"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/business"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/model"
	"html/template"
	"log"
	"net/http"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

var tpl *template.Template

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if err := tpl.ExecuteTemplate(res, "index.gohtml", u); err != nil {
		log.Println("Error en index:", err)
	}

}

func seccion(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	if err := tpl.ExecuteTemplate(res, "seccion.gohtml", u); err != nil {
		log.Println("Error en seccion:", err)
	}

}

func abmPersona(res http.ResponseWriter, req *http.Request) {
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	if err := tpl.ExecuteTemplate(res, "abmPersona.gohtml", nil); err != nil {
		log.Println("Error en abmPersona:", err)
	}
}

func altaUser(res http.ResponseWriter, req *http.Request) {
	if estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var u model.User

	if req.Method == http.MethodPost {

		usu := req.FormValue("usuario")
		pass := req.FormValue("password")
		nom := req.FormValue("nombre")
		ape := req.FormValue("apellido")

		u = business.CreaUsuario(usu, pass, nom, ape)
		if u.Error != nil {
			if err := tpl.ExecuteTemplate(res, "altaUser.gohtml", u); err != nil {
				log.Println("Error en altaUser:", err)
			}
			return
		}

	}

	if err := tpl.ExecuteTemplate(res, "altaUser.gohtml", u); err != nil {
		log.Println("Error en altaUser:", err)
	}
}

func login(res http.ResponseWriter, req *http.Request) {

	if estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	//TODO Revisar
	var u model.User
	if req.Method == http.MethodPost {
		usu := req.FormValue("usuario")
		pass := req.FormValue("password")
		u, ok := business.Autentica(usu, pass)
		if !ok {
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}
		log.Println("Login autentica: ", u)
		sID := uuid.New()
		c := &http.Cookie{
			Name:  sessionCookie,
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = usu

		if err := tpl.ExecuteTemplate(res, "index.gohtml", u); err != nil {
			log.Println("Error en index:", err)
		}
		return
	}

	if err := tpl.ExecuteTemplate(res, "login.gohtml", u); err != nil {
		log.Println("Error en login:", err)
	}
}

func logout(res http.ResponseWriter, req *http.Request) {

	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	c, _ := req.Cookie(sessionCookie)
	delete(dbSessions, c.Value)

	c = &http.Cookie{
		Name:   sessionCookie,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, c)

	http.Redirect(res, req, "/login", http.StatusSeeOther)

}
