package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	Apellido string
	Nombre   string
	Edad     int
}

type PorEdad []Persona

func (a PorEdad) Len() int           { return len(a) }
func (a PorEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

type PorApellido []Persona

func (a PorApellido) Len() int           { return len(a) }
func (a PorApellido) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorApellido) Less(i, j int) bool { return a[i].Apellido < a[j].Apellido }

func main() {

	p1 := Persona{
		Nombre:   "Fer",
		Apellido: "Czardas",
		Edad:     333,
	}

	p2 := Persona{
		Nombre:   "Raul",
		Apellido: "Barolo",
		Edad:     66,
	}

	p3 := Persona{
		Nombre:   "Aaul",
		Apellido: "Aarolo",
		Edad:     166,
	}

	e := []Persona{p1, p2, p3}
	fmt.Println(e)

	sort.Sort(PorEdad(e))

	fmt.Println(e)

	sort.Sort(PorApellido(e))

	fmt.Println(e)
}
