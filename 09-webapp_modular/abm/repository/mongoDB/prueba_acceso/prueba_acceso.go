package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mberliner/gobase/09-webapp_modular/abm/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx context.Context
)

func init() {

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://test:test987@localhost/go_test"))
	if err != nil {
		panic(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		panic(err)
	}
	col := client.Database("go_test").Collection("user")

	fmt.Println("Conectado a Mongodb: ", col)

	uC, e := getAll(col)
	if e != nil {
		panic(e)
	}
	fmt.Printf("Users: %T\n", uC)
	for i, u := range uC {
		fmt.Println(i, " - ", u)
	}

}

// Sólo para pruebas
func main() {
	fmt.Println("Conectado a Mongodb")
}

func getAll(col *mongo.Collection) ([]*model.User, error) {
	// passing bson.D{{}} matches all documents in the collection
	filter := bson.D{{}}
	var users []*model.User
	cur, err := col.Find(ctx, filter)
	if err != nil {
		return users, err
	}
	for cur.Next(ctx) {
		var t model.User
		err := cur.Decode(&t)
		if err != nil {
			return users, err
		}

		users = append(users, &t)
	}

	if err := cur.Err(); err != nil {
		return users, err
	}

	cur.Close(ctx)
	/*
		if len(users) == 0 {
			return users, mongo.ErrNoDocuments
		}
	*/
	return users, nil

}
