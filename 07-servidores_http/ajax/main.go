package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/asinc", asinc)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func asinc(res http.ResponseWriter, r *http.Request) {
	s := `Este texto llega de forma asincrónica`
	fmt.Fprintln(res, s)
}
