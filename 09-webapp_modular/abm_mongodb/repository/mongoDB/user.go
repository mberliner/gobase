package mongoDB

import (
	"context"
	"fmt"

	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Usuario  string             `json:"Usuario" bson:"usuario"`
	Nombre   string             `json:"Nombre" bson:"nombre"`
	Apellido string             `json:"Apellido" bson:"apellido"`
	Edad     string             `json:"Edad" bson:"edad"`
	Password string             `json:"Password" bson:"password"`
}

// UserRepository interface para acceso a datos de User
type UserRepository interface {
	Persiste(u *model.User) (*model.User, error)
	BuscaPorUsuario(usu string) ([]model.User, error)
}

// NewUserRepository para obtener repositorio de manera ordenada
func NewUserRepository(db *mongo.Database, ctx context.Context) UserRepository {
	return &userRepository{db, ctx}
}

type userRepository struct {
	db  *mongo.Database
	ctx context.Context
}

func (uR userRepository) Persiste(u *model.User) (*model.User, error) {

	uL := user{ID: primitive.NewObjectID(),
		Usuario:  u.Usuario,
		Nombre:   u.Nombre,
		Apellido: u.Apellido,
		Edad:     u.Edad,
		Password: u.Password,
	}

	_, err := uR.db.Collection("user").InsertOne(uR.ctx, uL)
	if err != nil {
		fmt.Println("Errrrr: ", err)
		return &model.User{}, err
	}

	return u, nil
}

func (uR userRepository) BuscaPorUsuario(usu string) ([]model.User, error) {

	cursor, err := uR.db.Collection("user").Find(uR.ctx, bson.M{"usuario": usu})
	if err != nil && err.Error() == "not found" {
		return nil, nil
	}
	if err != nil {
		return []model.User{}, err
	}

	suM := []model.User{}
	for cursor.Next(ctx) {
		var u user
		err := cursor.Decode(&u)
		if err != nil {
			return []model.User{}, err

		}
		uM := model.User{
			ID:       u.ID.String(),
			Usuario:  u.Usuario,
			Nombre:   u.Nombre,
			Apellido: u.Apellido,
			Edad:     u.Edad,
			Password: u.Password,
		}

		suM = append(suM, uM) //TODO usar model directo
	}

	return suM, nil
}
