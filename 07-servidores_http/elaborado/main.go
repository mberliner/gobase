package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
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

func graboCookie(res http.ResponseWriter, req *http.Request) {
	c := http.Cookie{
		Name:  "cookie-nombre",
		Value: "grabado-1",
	}

	//grabamos Cookie en navegador
	http.SetCookie(res, &c)
	io.WriteString(res, "Acabamos de grabar la cookie llamada: "+c.Name)
}

func leoCookie(res http.ResponseWriter, req *http.Request) {

	//grabamos Cookie en navegador
	c, err := req.Cookie("cookie-nombre")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Println("La Cookie leida es: ", c)
	io.WriteString(res, "Ya leimos su cookie gracias")
}

func grabaVisitas(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("cookie-visitas")
	//Si no existe la creo
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "cookie-visitas",
			Value: "0",
			Path:  "/",
		}
	}
	cont, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}
	cont++
	c.Value = strconv.Itoa(cont)

	//grabamos Cookie en navegador
	http.SetCookie(res, c)
	io.WriteString(res, "Visitas:"+c.Value)
}

func main() {

	http.HandleFunc("/", grabaVisitas)

	http.HandleFunc("/plantocookie", graboCookie)

	http.HandleFunc("/leocookie", leoCookie)

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
