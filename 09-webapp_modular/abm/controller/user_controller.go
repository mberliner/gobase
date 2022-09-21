package controller

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/mberliner/gobase/09-webapp_modular/abm/model"
	"github.com/mberliner/gobase/09-webapp_modular/abm/service"
)

func AltaUser(res http.ResponseWriter, req *http.Request) {
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

		u = service.UserS.CreaUsuario(usu, pass, nom, ape)

	}

	if err := tpl.ExecuteTemplate(res, "altaUser.gohtml", u); err != nil {
		log.Println("Error en altaUser:", err)
	}
}

func Login(res http.ResponseWriter, req *http.Request) {

	if estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	//TODO Revisar
	var u model.User
	if req.Method == http.MethodPost {
		usu := req.FormValue("usuario")
		pass := req.FormValue("password")
		u, ok := service.UserS.Autentica(usu, pass)
		if !ok {
			if err := tpl.ExecuteTemplate(res, "login.gohtml", u); err != nil {
				log.Println("Error en login:", err)
			}
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

func Logout(res http.ResponseWriter, req *http.Request) {

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
