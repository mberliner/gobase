package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func arch(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("template.gohtml"))

	var s string

	if req.Method == http.MethodPost {

		// open
		file, h, err := req.FormFile("archivo")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		fmt.Println("\nArchivo:", file, "\nheader:", h, "\nerr", err)

		// read
		bs, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(res, "template.gohtml", s)
	if err != nil {
		log.Fatalln("error executing template", err)
	}

}

func main() {

	http.HandleFunc("/archivos", arch)

	//Por si no hay favicon hay browsers que lo piden una y otra vez
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}
