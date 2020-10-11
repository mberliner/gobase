package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/repository"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
var dbUsers = map[string]repository.User{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/seccion", seccion)
	http.HandleFunc("/altausuario", altaUser)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	//	http.HandleFunc("/persona/alta", altaPersona)
	http.HandleFunc("/abmPersona", abmPersona)
	//	http.HandleFunc("/persona/baja", bajaPersona)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func seccion(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "seccion.gohtml", u)
}

func abmPersona(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "abmPersona.gohtml", u)
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

		sU, err := repository.UR.BuscaPorUsuario(usu)
		if err != nil || len(sU) > 0 {
			http.Error(res, "El usuario ya existe, elija otro nombre", http.StatusForbidden)
			return
		}

		sID := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = usu

		encrPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "No puedo encriptar pass", http.StatusInternalServerError)
			return
		}

		u := repository.User{Usuario: usu, Nombre: nom, Apellido: ape, Password: encrPass}
		_, err = repository.UR.Persiste(u)
		if err != nil {
			fmt.Println("err persiste:", err)
			http.Error(res, "No persiste", http.StatusInternalServerError)
			return
		}

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "altaUser.gohtml", u)
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
		sU, err := repository.UR.BuscaPorUsuario(usu)
		if err != nil || len(sU) == 0 {
			fmt.Println("user:------------------entra2-------------------------", sU, err)
			http.Error(res, "Usuario o Password inválidos", http.StatusForbidden)
			return
		}

		err = bcrypt.CompareHashAndPassword(sU[0].Password, []byte(pass))
		if err != nil {
			fmt.Println("user:------------------entra3-------------------------", sU, err, pass, sU[0].Password)
			http.Error(res, "Usuario o Password inválidos", http.StatusForbidden)
			return
		}

		sID := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = usu

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "login.gohtml", sU[0])
}

func logout(res http.ResponseWriter, req *http.Request) {

	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	c, _ := req.Cookie("session")
	delete(dbSessions, c.Value)

	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, c)

	http.Redirect(res, req, "/login", http.StatusSeeOther)

}
