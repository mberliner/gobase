package repository

import (
	"fmt"

	"github.com/globalsign/mgo"
)

var db *mgo.Database
var UserRepo UserRepository

var PersonaRepo PersonaRepository

func init() {

	var err error

	s, err := mgo.Dial("mongodb://test:test987@localhost/go_test")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	db = s.DB("go_test")

	fmt.Println("Conectado a Mongodb")

	UserRepo = NewUserRepository(db)
	PersonaRepo = NewPersonaRepository(db)

}
