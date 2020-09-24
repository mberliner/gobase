package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func formateaFecha(t time.Time) string {
	return t.Format("02-01-2006") //02=Dia 01=Mes 2006=Año
}

func main() {

	//FuncMap se usa para pasar funciones a templates
	var fm = template.FuncMap{
		"formatea": formateaFecha,
	}
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("template.txt"))

	type persona struct {
		Apellido string
		Nombre   string
		Edad     int
	}

	p1 := persona{
		Nombre:   "Fer",
		Apellido: "Czardas",
		Edad:     333,
	}

	p2 := persona{
		Nombre:   "Raul",
		Apellido: "Barolo",
		Edad:     66,
	}

	p3 := persona{
		Nombre:   "Aaul",
		Apellido: "Aarolo",
		Edad:     166,
	}

	personas := []persona{p1, p2, p3}
	numeros := []int{100, 10000, 37}
	mapa := map[string]string{
		"Key1": "Elemento1",
		"Key2": "Elemento2",
	}

	//Paso un sólo dato pero puede agrupar todo lo que necesito
	data := struct {
		Persona []persona
		Num     []int
		Cosas   map[string]string
		Fecha   time.Time
	}{
		personas,
		numeros,
		mapa,
		time.Now(),
	}

	//sin funciones
	//err := tpl.Execute(os.Stdout, data)

	err := tpl.ExecuteTemplate(os.Stdout, "template.txt", data)
	if err != nil {
		log.Fatalln(err)
	}
}
