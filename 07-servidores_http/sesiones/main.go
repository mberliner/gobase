package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	Usuario  string
	Nombre   string
	Apellido string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/seccion", seccion)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("session")
	if err != nil {
		//Si no existe la cookie de sesion creo una nueva
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
	}

	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("usuario")
		f := req.FormValue("nombre")
		l := req.FormValue("apellido")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func seccion(res http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(res, "seccion.gohtml", u)
}
