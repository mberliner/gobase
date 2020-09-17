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

type person struct {
	nombre   string
	apellido string
	direccion
}

type person1 struct {
	nombre   string
	apellido string
}

func main() {

	persona1 := person1{"Fer", "Jones"} // Asigna por el orden
	persona2 := person{nombre: "Sergio",
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

	fmt.Println(persona1)

	fmt.Println("1---------------------------")
	persona2.apellido = "Rober"
	persona2.imprimir()

	fmt.Println("2---------------------------")
	persona2.actualizarApe("Xxxxxxxxxxxxxxxx") //El compilado hace la traducción entre variable y su puntero dentro de la función
	persona2.imprimir()

	//imprime c/u y los datos internos se promueven a la estructura exterior
	fmt.Println("Imprimo valores internos: ", persona2.calle, persona2.numero)

	fmt.Println("3 struct anonima---------------------------")

	fmt.Println(persona3)

}

func (p person) imprimir() {
	fmt.Printf("Es una persona: %+v\n", p)
}

func (p *person) actualizarApe(apellido string) {
	p.apellido = apellido
}
