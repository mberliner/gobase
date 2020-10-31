package repository

import (
	"strings"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/model"
)

type persona struct {
	ID              bson.ObjectId //`json:"Id" bson:"Id"`
	Nombre          string        //`json:"Nombre" bson:"Nombre"`
	Apellido        string        //`json:"Apellido" bson:"Apellido"`
	FechaNacimiento string        //`json:"FechaNacimiento" bson:"FechaNacimiento"`
}

type PersonaRepository interface {
	Persiste(p *model.Persona) (*model.Persona, error)
	Borra(id string) error
	BuscaTodo() ([]model.Persona, error)
	BuscaPorID(id string) (*model.Persona, error)
	Actualiza(p *model.Persona) (*model.Persona, error)
}

type personaRepository struct {
	db *mgo.Database
}

func NewPersonaRepository(db *mgo.Database) *personaRepository {
	return &personaRepository{db}
}

func (pR personaRepository) Persiste(p *model.Persona) (*model.Persona, error) {

	pL := persona{ID: bson.NewObjectId(),
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: p.FechaNacimiento,
	}
	err := db.C("persona").Insert(pL)
	if err != nil {
		return &model.Persona{}, err
	}

	return p, nil
}

func (pR personaRepository) Borra(id string) error {

	id1 := bson.ObjectIdHex(id)
	err := db.C("persona").Remove(bson.M{"id": id1})
	if err != nil {
		return err
	}

	return nil
}

func (pR personaRepository) BuscaTodo() ([]model.Persona, error) {

	personas := []persona{}
	err := db.C("persona").Find(bson.M{}).All(&personas)
	if err != nil {
		return []model.Persona{}, err
	}

	var rP []model.Persona
	for _, p := range personas {

		idd := strings.Split(p.ID.String(), `"`)

		per := model.Persona{ID: idd[1],
			Nombre:          p.Nombre,
			Apellido:        p.Apellido,
			FechaNacimiento: p.FechaNacimiento,
		}

		rP = append(rP, per)
	}

	return rP, nil
}

func (pR personaRepository) BuscaPorID(id string) (*model.Persona, error) {
	p := persona{}
	id1 := bson.ObjectIdHex(id)

	err := db.C("persona").Find(bson.M{"id": id1}).One(&p)
	if err != nil && err.Error() == "not found" {
		return &model.Persona{}, nil
	}
	if err != nil {
		return &model.Persona{}, err
	}
	idd := strings.Split(p.ID.String(), `"`)
	perM := model.Persona{ID: idd[1],
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: p.FechaNacimiento,
	}
	return &perM, nil
}

func (pR personaRepository) Actualiza(p *model.Persona) (*model.Persona, error) {

	pL := persona{ID: bson.ObjectIdHex(p.ID),
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: p.FechaNacimiento,
	}

	err := db.C("persona").Update(bson.M{"id": pL.ID}, &pL)
	if err != nil {
		return &model.Persona{}, err
	}

	return p, nil
}
