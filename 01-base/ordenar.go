package main

import (
	"fmt"
	"sort"
)

type persona struct {
	Apellido string
	Nombre   string
	Edad     int
}

type porEdad []persona

func (a porEdad) Len() int           { return len(a) }
func (a porEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a porEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

type porApellido []persona

func (a porApellido) Len() int           { return len(a) }
func (a porApellido) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a porApellido) Less(i, j int) bool { return a[i].Apellido < a[j].Apellido }

func main() {

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

	e := []persona{p1, p2, p3}
	fmt.Println(e)

	sort.Sort(porEdad(e))

	fmt.Println(e)

	sort.Sort(porApellido(e))

	fmt.Println(e)
}
