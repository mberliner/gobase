package service

import (
	"log"

	"github.com/mberliner/gobase/09-webapp_modular/abm/model"
	"github.com/mberliner/gobase/09-webapp_modular/abm/repository/mongoDB"
)

type personaService struct {
	personaRepo mongoDB.PersonaRepository
}

// PersonaService interface para poder realizar tests del negocio de persona
type PersonaService interface {
	CreaPersona(nom string, ape string, fechaNacimiento string) model.Personas
	BuscaTodo() model.Personas
	BorraPersona(id string) model.Personas
	BuscaPersona(id string) model.Personas
	ActualizaPersona(id string, nom string, ape string, fechaNacimiento string) model.Personas
}

// NewPersonaService para obtener megocio de forma ordenada
func NewPersonaService(pR mongoDB.PersonaRepository) PersonaService {
	return &personaService{pR}
}

func (pS personaService) CreaPersona(nom string, ape string, fechaNacimiento string) model.Personas {

	p := &model.Persona{Nombre: nom, Apellido: ape, FechaNacimiento: fechaNacimiento}
	p, err := pS.personaRepo.Persiste(p)
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

func (pS personaService) BuscaTodo() model.Personas {

	ps, err := pS.personaRepo.BuscaTodo()
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

func (pS personaService) BorraPersona(id string) model.Personas {

	err := pS.personaRepo.Borra(id)
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

func (pS personaService) BuscaPersona(id string) model.Personas {

	p, err := pS.personaRepo.BuscaPorID(id)
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

func (pS personaService) ActualizaPersona(id string, nom string, ape string, fechaNacimiento string) model.Personas {

	p := &model.Persona{Nombre: nom, Apellido: ape, FechaNacimiento: fechaNacimiento, ID: id}
	p, err := pS.personaRepo.Actualiza(p)
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
