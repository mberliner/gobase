package repository

import (
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type persona struct {
	ID              bson.ObjectId //`json:"ID" bson:"_id"`
	Nombre          string //`json:"Nombre" bson:"Nombre"`
	Apellido        string //`json:"Apellido" bson:"Apellido"`
	FechaNacimiento string //`json:"FechaNacimiento" bson:"FechaNacimiento"`
}

type PersonaRepository struct {
	db  *mgo.Database
}

func NewPersonaRepository(db  *mgo.Database) *PersonaRepository {
	return &PersonaRepository{db}
}

func (pR PersonaRepository) Persiste(p model.Persona) (model.Persona, error) {

log.Println("ObjectId:", bson.NewObjectId())

	pL := persona{ID: bson.NewObjectId(),
		Nombre:   p.Nombre,
		Apellido: p.Apellido,
		FechaNacimiento: p.FechaNacimiento,
	}
	err := db.C("persona").Insert(pL)
	if err != nil {
		return model.Persona{}, err
	}

	return p, nil
}

func (pR PersonaRepository) Borra(id string) error {

	err := db.C("persona").Remove(bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

func (pR PersonaRepository) BuscaTodo() ([]model.Persona, error) {

	personas := []persona{}
	err := db.C("persona").Find(bson.M{}).All(&personas)
	if err != nil {
		return []model.Persona{}, err
	}

	var rP []model.Persona
	for _, p := range personas {
		per := model.Persona{ID: p.ID.String(),
			Nombre:          p.Nombre,
			Apellido:        p.Apellido,
			FechaNacimiento: p.FechaNacimiento,
		}

		rP = append(rP, per)
	}

	return rP, nil
}

func (pR PersonaRepository) BuscaPorId(id string) (model.Persona, error) {
	p := persona{}
	err := db.C("persona").Find(bson.M{"ID": id}).One(&p)
	if err != nil && err.Error() == "not found" {
		return model.Persona{}, nil
	}
	if err != nil {
		return model.Persona{}, err
	}

	perM := model.Persona{ID: p.ID.String(),
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: p.FechaNacimiento,
	}
	return perM, nil
}

func (pR PersonaRepository) Actualiza(p model.Persona) (model.Persona, error) {

	pL := persona{ID: bson.ObjectId(p.ID),
			Nombre: p.Nombre,
			Apellido: p.Apellido,
			FechaNacimiento: p.FechaNacimiento,
			}

	err := db.C("persona").Update(bson.M{"ID": p.ID}, &pL)
	if err != nil {
		return model.Persona{}, err
	}


	return p, nil
}
