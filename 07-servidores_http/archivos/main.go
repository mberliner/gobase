package main

import (
	_ "fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func arch(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("template.gohtml"))

	var texto string

	if req.Method == http.MethodPost {

		//Abro archivo del post
		file, head, err := req.FormFile("archivo")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		//		fmt.Println("\nArchivo:", file, "\nheader:", head, "\nerr", err)

		// leo datos
		bs, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		texto = string(bs)

		// guardo archivo
		destino, err := os.Create(filepath.Join("./archivo_destino/", head.Filename))
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer destino.Close()

		_, err = destino.Write(bs)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(res, "template.gohtml", texto)
	if err != nil {
		log.Fatalln("Error al ejecutar template", err)
	}

}

func main() {

	http.HandleFunc("/archivos", arch)

	//Por si no hay favicon hay browsers que lo piden una y otra vez
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}
