package main

import (
	"log"
	"os"
	"text/template"
)

type persona struct {
	Apellido  string
	Nombre    string
	Edad      float64
	Esperanza float64
}

//TimeToDie se exporta para poder ser usanda en templates
func (p persona) TimeToDie() float64 {
	return (p.Esperanza - p.Edad)
}

//TimeToDieAvg se exporta para poder ser usanda en templates
func (p persona) TimeToDieAvg(v float64) float64 {
	return (p.Esperanza - p.Edad) / v
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.txt"))

}

func main() {

	p1 := persona{
		Nombre:    "Fer",
		Apellido:  "Czardas",
		Edad:      33,
		Esperanza: 87.6,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
