package controller

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func init() {
	var fm = template.FuncMap{
		"formateaFecha": formateaFecha,
	}

	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

var tpl *template.Template

func formateaFecha(t time.Time) string {
	return t.Format("02-01-2006") //02=Dia 01=Mes 2006=Año
}

func Index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if err := tpl.ExecuteTemplate(res, "index.gohtml", u); err != nil {
		log.Println("Error en index:", err)
	}

}

func Seccion(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !estaLogueado(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	if err := tpl.ExecuteTemplate(res, "seccion.gohtml", u); err != nil {
		log.Println("Error en seccion:", err)
	}

}
