package mongoDB

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (

	//UserRepo Repositorio para manejo de acceso a datos de usuario
	UserRepo UserRepository
	//PersonaRepo Repositorio para manejo de acceso a datos de persona
	PersonaRepo PersonaRepository
	ctx         context.Context
)

func init() {

	var err error
	fmt.Println("Inicio MongoDB, nuevo driver")
	//TODO ver timeout
	ctx, _ := context.WithTimeout(context.Background(), 720*time.Second)
	//	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	//	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://test:test987@localhost/go_test"))
	if err != nil {
		panic(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		panic(err)
	}
	db := client.Database("go_test")

	fmt.Println("Conectado a Mongodb: ", db)

	UserRepo = NewUserRepository(db, ctx)
	PersonaRepo = NewPersonaRepository(db, ctx)

}
