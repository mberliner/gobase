package repository

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var db *mgo.Database
var UR *UserRepository

var PR *PersonaRepository

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

	UR = NewUserRepository(db)
	PR = NewPersonaRepository(db)

}
