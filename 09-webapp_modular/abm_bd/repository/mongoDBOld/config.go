package mongoDBOld

import (
	"fmt"

	"github.com/globalsign/mgo"
)

var (
	//Lamentablemente el driver de mongo no posee interfaces de forma que no podemos testear el repo de manera unitaria
	//Para hacer tal cosa deberíamos generar un wrapper para cada objeto del driver de mgo y usarlo en nuestro desarrollo
	//Esto hace que agregemos una capa más (agregando más posibles errores) con el único fin de test
	//No vale la pena el esfuerzo, es preferible hacer test de integración/componentes con una BD Mongo real o en memoria
	db *mgo.Database
	//UserRepo Repositorio para manejo de acceso a datos de usuario
	UserRepo UserRepository
	//PersonaRepo Repositorio para manejo de acceso a datos de persona
	PersonaRepo PersonaRepository
)

func init() {

	var err error
	fmt.Println("Inicio MongoDB, antiguo driver")
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
