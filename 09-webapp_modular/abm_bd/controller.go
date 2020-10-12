package main

import (
	"github.com/google/uuid"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/business"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/repository"
	"html/template"
	"net/http"
	"log"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

var tpl *template.Template

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if err:= tpl.ExecuteTemplate(res, "index.gohtml", u); err!=nil {
		log.Println("Error en index:", err)
	}

}

func seccion(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	if err:= tpl.ExecuteTemplate(res, "seccion.gohtml", u); err!=nil {
		log.Println("Error en seccion:", err)
	}
	
}

func abmPersona(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	if err:= tpl.ExecuteTemplate(res, "abmPersona.gohtml", u); err!=nil {
		log.Println("Error en abmPersona:", err)
	}
}

func altaUser(res http.ResponseWriter, req *http.Request) {
	if estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var u repository.User

	if req.Method == http.MethodPost {

		usu := req.FormValue("usuario")
		pass := req.FormValue("password")
		nom := req.FormValue("nombre")
		ape := req.FormValue("apellido")

		_, err := business.CreaUsuario(usu, pass, nom, ape)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}

	if err:= tpl.ExecuteTemplate(res, "altaUser.gohtml", u); err!=nil {
		log.Println("Error en altaUser:", err)
	}
}

func login(res http.ResponseWriter, req *http.Request) {

	if estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	//TODO Revisar
	var sU = make([]repository.User, 1)
	if req.Method == http.MethodPost {
		usu := req.FormValue("usuario")
		pass := req.FormValue("password")

		if !business.Autentica(usu, pass) {
			http.Error(res, "Usuario o Password inválidos", http.StatusForbidden)
			return
		}

		sID := uuid.New()
		c := &http.Cookie{
			Name:  sessionCookie,
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = usu

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if err:= tpl.ExecuteTemplate(res, "login.gohtml", sU[0]); err!=nil {
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
