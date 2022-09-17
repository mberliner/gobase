package main

import (
	"fmt"
)

func main() {

	var p persona

	//TODO ver un lugar mas elegante para ubicarlo
	defer finalizaRepo()

	personaRepo.crea()

	p.nombre = "George"
	p.apellido = "Sansonin"
	p.edad = 39
	p.id = 1
	p1 := personaRepo.persiste(p)

	p.nombre = "Falcon"
	p.apellido = "Situs"
	p.edad = 23
	p.id = 2
	p2 := personaRepo.persiste(p)

	p.nombre = "Manual"
	p.apellido = "Onargleb"
	p.edad = 33
	p.id = 3
	personaRepo.persiste(p)

	personas := personaRepo.traeTodas()
	fmt.Println("traje Todas las Personas", personas)

	p1.nombre = "Jorge"
	personaRepo.actualiza(p1)

	fmt.Println("Voy a traer una persona id 1")
	px := personaRepo.findByID(1)
	fmt.Println("P1:", px)

	fmt.Println("Voy a hacer una transación con:", p2)
	personaRepo.operacionesComplejas(p2)
	fmt.Println("Resultado:", personaRepo.traeTodas())

}
