package main

import (
	"fmt"
)

type direccion struct {
	calle  string
	numero int
	piso   int
	dpto   string
}

type persona struct {
	nombre   string
	apellido string
	direccion
}

type personaSimple struct {
	nombre   string
	apellido string
}

func main() {

	persona1 := personaSimple{"Fer", "Jones"} // Asigna por el orden
	persona2 := persona{nombre: "Sergio",
		apellido: "Renen",
		direccion: direccion{
			calle:  "Libertad",
			numero: 100,
			piso:   245,
			dpto:   "B"},
	}

	//Anonim estructura, otra forma definido en la misma creación del valor...
	persona3 := struct {
		nombre   string
		apellido string
	}{nombre: "Vixento",
		apellido: "Valvolatti",
	}

	fmt.Printf("Simple: %T - %v\n", persona1, persona1)

	fmt.Println("1---------------------------")
	persona2.apellido = "Rober"
	persona2.imprimir()

	fmt.Println("2---------------------------")
	persona2.actualizarApe("Xxxxxxxxx") //El compilado hace la traducción entre variable y su puntero dentro de la función
	persona2.imprimir()

	//imprime c/u y los datos internos se promueven a la estructura exterior
	fmt.Println("Imprimo algunos valores internos promovidos: ", persona2.calle, persona2.numero)

	fmt.Println("3 struct anonima---------------------------")

	fmt.Printf("Simple: %T - %v\n", persona3, persona3)

}

func (p persona) imprimir() {
	fmt.Printf("Es una persona:  %T - %+v\n", p, p)
}

func (p *persona) actualizarApe(apellido string) {
	p.apellido = apellido
}
