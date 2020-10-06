package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var p persona

	db := iniciaBD()
	defer db.Close()
	pRepo := newPersonaRepository(db)

	pRepo.crea()

	p.nombre = "George"
	p.apellido = "Sansonin"
	p.edad = 39
	p.id = 1
	p1 := pRepo.persiste(p)

	p.nombre = "Falcon"
	p.apellido = "Situs"
	p.edad = 23
	p.id = 2
	p2 := pRepo.persiste(p)

	p.nombre = "Manual"
	p.apellido = "Onargleb"
	p.edad = 33
	p.id = 3
	pRepo.persiste(p)

	personas := pRepo.presentaTodo()
	fmt.Println("traje Todas las Personas", personas)

	p1.nombre = "Jorge"
	pRepo.actualiza(p1)

	fmt.Println("Voy a traer una persona id 1")
	px := pRepo.findByID(1)
	fmt.Println("P1:", px)

	fmt.Println("Voy a hacer una transación con:", p2)
	pRepo.operacionesComplejas(p2)
	fmt.Println("Resultado:", pRepo.presentaTodo())

}

func iniciaBD() *sql.DB {
	db, err := sql.Open("mysql", "root:soygroot@tcp(localhost:3306)/go_test?charset=utf8")
	check(err)

	err = db.Ping() //Veo si accedo bien
	check(err)

	return db
}
