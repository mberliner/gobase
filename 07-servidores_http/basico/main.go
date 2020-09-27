package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type myHandler int

func (m myHandler) ServeHTTP(resW http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	//Seteo variables del header
	resW.Header().Set("My-Key", "Datos de la clave")
	resW.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Datos que van en la respuesta
	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}
	tpl.ExecuteTemplate(resW, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var mh myHandler
	http.ListenAndServe(":8080", mh)
}
