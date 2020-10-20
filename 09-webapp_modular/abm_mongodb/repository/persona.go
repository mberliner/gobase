package repository

import (
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strings"
)

type persona struct {
	Id              bson.ObjectId //`json:"Id" bson:"_Id"`
	Nombre          string        //`json:"Nombre" bson:"Nombre"`
	Apellido        string        //`json:"Apellido" bson:"Apellido"`
	FechaNacimiento string        //`json:"FechaNacimiento" bson:"FechaNacimiento"`
}

type PersonaRepository struct {
	db *mgo.Database
}

func NewPersonaRepository(db *mgo.Database) *PersonaRepository {
	return &PersonaRepository{db}
}

func (pR PersonaRepository) Persiste(p model.Persona) (model.Persona, error) {

	pL := persona{Id: bson.NewObjectId(),
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: p.FechaNacimiento,
	}
	err := db.C("persona").Insert(pL)
	if err != nil {
		return model.Persona{}, err
	}

	return p, nil
}

func (pR PersonaRepository) Borra(id string) error {

	id1 := bson.ObjectIdHex(id)
	err := db.C("persona").Remove(bson.M{"id": id1})
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

		log.Println("rangeo:", p)

		idd := strings.Split(p.Id.String(), `"`)

		per := model.Persona{ID: idd[1],
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
	id1 := bson.ObjectIdHex(id)

	err := db.C("persona").Find(bson.M{"id": id1}).One(&p)
	log.Println("busca por ID:", p, err)
	if err != nil && err.Error() == "not found" {
		return model.Persona{}, nil
	}
	if err != nil {
		return model.Persona{}, err
	}
	idd := strings.Split(p.Id.String(), `"`)
	perM := model.Persona{ID: idd[1],
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: p.FechaNacimiento,
	}
	return perM, nil
}

func (pR PersonaRepository) Actualiza(p model.Persona) (model.Persona, error) {

	pL := persona{Id: bson.ObjectIdHex(p.ID),
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: p.FechaNacimiento,
	}

	err := db.C("persona").Update(bson.M{"id": pL.Id}, &pL)
	if err != nil {
		return model.Persona{}, err
	}

	return p, nil
}
