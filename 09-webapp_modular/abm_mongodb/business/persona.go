package business

import (
	"log"

	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/model"
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/repository"
)

type personaBusiness struct {
	personaRepo repository.PersonaRepository
}

type PersonaBusiness interface {
	CreaPersona(nom string, ape string, fechaNacimiento string) model.Personas
	BuscaTodo() model.Personas
	BorraPersona(id string) model.Personas
	BuscaPersona(id string) model.Personas
	ActualizaPersona(id string, nom string, ape string, fechaNacimiento string) model.Personas
}

func NewPersonaBusiness(pR repository.PersonaRepository) PersonaBusiness {
	return &personaBusiness{pR}
}
func (pB personaBusiness) CreaPersona(nom string, ape string, fechaNacimiento string) model.Personas {

	p := &model.Persona{Nombre: nom, Apellido: ape, FechaNacimiento: fechaNacimiento}
	p, err := pB.personaRepo.Persiste(p)
	if err != nil {
		log.Println("Error persiste Pesona:", err)
		mP := model.Personas{}
		mP.Error = err
		return mP
	}

	personas := []model.Persona{*p}

	mP := model.Personas{
		PersonasM: personas,
		Error:     nil,
		Mensaje:   "Persona Creada Ok",
	}
	return mP
}

func (pB personaBusiness) BuscaTodo() model.Personas {

	ps, err := pB.personaRepo.BuscaTodo()
	if err != nil {
		log.Println("Error buscaTodo:", err)
		mP := model.Personas{}
		mP.Error = err
		return mP
	}

	mP := model.Personas{
		PersonasM: ps,
		Error:     nil,
		Mensaje:   "Carga ok",
	}

	return mP
}

func (pB personaBusiness) BorraPersona(id string) model.Personas {

	err := pB.personaRepo.Borra(id)
	if err != nil {
		log.Println("Error borraPersona:", err)
		mP := model.Personas{}
		mP.Error = err
		return mP
	}
	personas := []model.Persona{}
	per := model.Persona{ID: id,
		Nombre:   "",
		Apellido: "",
	}

	personas = append(personas, per)

	mP := model.Personas{
		PersonasM: personas,
		Error:     nil,
		Mensaje:   "Borrado ok",
	}

	return mP
}

func (pB personaBusiness) BuscaPersona(id string) model.Personas {

	p, err := pB.personaRepo.BuscaPorID(id)
	if err != nil {
		log.Println("Error buscaPersona:", err)
		mP := model.Personas{}
		mP.Error = err
		return mP
	}

	mP := model.Personas{
		PersonasM: []model.Persona{*p},
		Error:     nil,
		Mensaje:   "Busca ok",
	}

	return mP
}

func (pB personaBusiness) ActualizaPersona(id string, nom string, ape string, fechaNacimiento string) model.Personas {

	p := &model.Persona{Nombre: nom, Apellido: ape, FechaNacimiento: fechaNacimiento, ID: id}
	p, err := pB.personaRepo.Actualiza(p)
	log.Println("Error actualiza Persona con fecha:", p, err)
	if err != nil {
		log.Println("Error actualiza Pesona:", err)
		mP := model.Personas{}
		mP.Error = err
		return mP
	}

	mP := model.Personas{
		PersonasM: []model.Persona{*p},
		Error:     nil,
		Mensaje:   "Persona Actualizada Ok",
	}
	return mP
}
