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

	personas := p.presentaTodo(db)
	fmt.Println("traje Todas las Personas", personas)

	fmt.Println("Voy a traer una persona id 1")
	p1 := p.traePersonaPorID(db, 1)
	fmt.Println("P1:", p1)

	p.nombre = "George"
	p.apellido = "Sansonin"
	p.edad = 39
	p.id = 3
	p2 := p.persiste(db)

	fmt.Println("P2:", p2)
	p2.operacionesComplejas(db)
	fmt.Println("P2:", p.traePersonaPorID(db, 3))
}

func iniciaBD() *sql.DB {
	db, err := sql.Open("mysql", "root:soygroot@tcp(localhost:3306)/go_test?charset=utf8")
	check(err)

	err = db.Ping() //Veo si accedo bien
	check(err)

	return db
}
