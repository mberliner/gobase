package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func mineral(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("template.gohtml"))

	//Tomo param del form (al igual que de una url)
	v := req.FormValue("nombre")

	err := tpl.ExecuteTemplate(res, "template.gohtml", "Afortunado Mineral!   "+v)
	if err != nil {
		log.Fatalln("error executing template", err)
	}

}

func animal(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "esta es la info de animal")
}

func persona(res http.ResponseWriter, req *http.Request) {

	//Tomar parametros de URL
	v := req.FormValue("q")
	v1 := req.FormValue("w")
	io.WriteString(res, "Esta es la info de persona, si hay parámetros sale acá: "+v+"  "+v1)
}

//Sirve un archivo
func file(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "file/dan.jpg")
}

func main() {

	//Con barra final toma todo path porterior
	http.HandleFunc("/persona/", persona)

	//sin barra es estricto
	http.HandleFunc("/animal", animal)

	http.HandleFunc("/mineral", mineral)

	//Archivo
	http.HandleFunc("/file/dan.jpg", file)

	//Sirve archivos en el dir file (aunque llegue una urls con /filesx
	http.Handle("/filesx/", http.StripPrefix("/filesx/", http.FileServer(http.Dir("file"))))

	//Por si no hay favicon hay browsers que lo piden una y otra vez
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}
