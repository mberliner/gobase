package main

import (
	"fmt"
)

//declaraciones siempre con el mismo formato "type identificador tipo"
type agente struct {
	nombre   string
	apellido string
	licencia int
}

type cirujano struct {
	nombre    string
	apellido  string
	matricula int
}

//Damos comportamiento a los tipos
//Usado para lograr polimorfismo
type humano interface {
	hablar()
}

func main() {
	juan := cirujano{
		nombre:    "Juan",
		apellido:  "Rivas",
		matricula: 254321,
	}
	james := agente{
		nombre:   "James",
		apellido: "Bond",
		licencia: 007,
	}

	//Usamos comportamiento
	juan.hablar()
	james.hablar()

	//Es posible usar un puntero si el receiver es un tipo (al revés no es posible)
	//Esto es porque asi están definidos los "method sets"
	pJames := &james
	fmt.Println("Ejemplo Probando Method Sets. Llamo con puntero: ")
	pJames.hablar()

	//Llamamos desde afuera
	fmt.Println("\nEjemplo Polimorfismo--------------------------")
	quienHabla(juan)
	quienHabla(james)

}

//Solo por implementar este método ya implementan la interface humano
func (a agente) hablar() {
	fmt.Printf("Soy un %T\n", a)
	fmt.Println("Me gusta tirar tiros")
}

func (c cirujano) hablar() {
	fmt.Printf("Soy un %T\n", c)
	fmt.Println("Me gusta Operar")
}

func quienHabla(h humano) {
	fmt.Printf("Soy un ---- %T\n", h)
	fmt.Println(h, "Voy a hablar")
	h.hablar() //Esto es polimorfismo, puedo llamar a los humanos que quieran hablar incluso si hoy todavia no existen
}
