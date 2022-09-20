package mongoDB

import (
	"context"
	"fmt"
	"strings"

	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type persona struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	Nombre          string             `json:"Nombre" bson:"nombre"`
	Apellido        string             `json:"Apellido" bson:"apellido"`
	FechaNacimiento string             `json:"FechaNacimiento" bson:"fechaNacimiento"`
}

// PersonaRepository interface para manejo de acceso a datos de Persona
type PersonaRepository interface {
	Persiste(p *model.Persona) (*model.Persona, error)
	Borra(id string) error
	BuscaTodo() ([]model.Persona, error)
	BuscaPorID(id string) (*model.Persona, error)
	Actualiza(p *model.Persona) (*model.Persona, error)
}

// NewPersonaRepository para obtener repositorio de manera ordenada
func NewPersonaRepository(db *mongo.Database, ctx context.Context) PersonaRepository {
	return &personaRepository{db, ctx}
}

type personaRepository struct {
	db  *mongo.Database
	ctx context.Context
}

func (pR personaRepository) Persiste(p *model.Persona) (*model.Persona, error) {

	pL := persona{ID: primitive.NewObjectID(),
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: p.FechaNacimiento,
	}
	res, err := pR.db.Collection("persona").InsertOne(pR.ctx, pL)
	if err != nil {
		return &model.Persona{}, err
	}
	p.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return p, nil
}

func (pR personaRepository) Borra(id string) error {

	id1, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	res, err := pR.db.Collection("persona").DeleteOne(pR.ctx, bson.M{"_id": id1})
	if err != nil {
		return err
	}
	fmt.Println("Cant. Borrado: ", res.DeletedCount)

	return nil
}

func (pR personaRepository) BuscaTodo() ([]model.Persona, error) {

	cursor, err := pR.db.Collection("persona").Find(pR.ctx, bson.M{})
	if err != nil {
		return []model.Persona{}, err
	}

	rP := []model.Persona{}
	for cursor.Next(ctx) {
		var p persona
		err := cursor.Decode(&p)
		if err != nil {
			return []model.Persona{}, err
		}

		idd := strings.Split(p.ID.String(), `"`)
		pM := model.Persona{ID: idd[1],
			Nombre:          p.Nombre,
			Apellido:        p.Apellido,
			FechaNacimiento: p.FechaNacimiento,
		}
		rP = append(rP, pM)
	}

	return rP, nil
}

func (pR personaRepository) BuscaPorID(id string) (*model.Persona, error) {
	p := persona{}
	id1, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &model.Persona{}, err
	}
	err = pR.db.Collection("persona").FindOne(pR.ctx, bson.M{"_id": id1}).Decode(&p)
	if err != nil && err.Error() == "not found" {
		return &model.Persona{}, nil
	}
	if err != nil {
		return &model.Persona{}, err
	}

	perM := model.Persona{ID: p.ID.Hex(),
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: p.FechaNacimiento,
	}
	return &perM, nil
}

func (pR personaRepository) Actualiza(p *model.Persona) (*model.Persona, error) {

	id1, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		return &model.Persona{}, err
	}

	filter := bson.M{"_id": bson.M{"$eq": id1}}
	actualiza := bson.M{"$set": bson.M{"nombre": p.Nombre,
		"apellido":        p.Apellido,
		"fechaNacimiento": p.FechaNacimiento}}

	res, err := pR.db.Collection("persona").UpdateOne(ctx, filter, actualiza)
	if err != nil {
		return &model.Persona{}, err
	}
	fmt.Println("Cant. Actualizado: ", res.ModifiedCount)

	return p, nil

}
