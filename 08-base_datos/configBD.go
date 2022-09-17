package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var personaRepo *personaRepository
var db *sql.DB

func init() {
	db = iniciaRepo()

	personaRepo = newPersonaRepository(db)

}

func iniciaRepo() *sql.DB {
	db, err := sql.Open("mysql", "test:test987@tcp(localhost:3306)/go_test?charset=utf8")
	check(err)
	err = db.Ping() //Veo si accedo bien
	check(err)

	return db
}

func finalizaRepo() {
	db.Close()
}
