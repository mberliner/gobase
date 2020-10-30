package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //Sólo para iniciar database/sql
)

var db *sql.DB
var UserRepo UserRepository
var PersonaRepo PersonaRepository

func init() {
	var err error
	db, err = sql.Open("mysql", "test:test987@tcp(localhost:3306)/go_test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Conectado a Mysql")

	UserRepo = NewUserRepository(db)
	PersonaRepo = NewPersonaRepository(db)

}
