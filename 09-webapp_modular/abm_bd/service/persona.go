package service

import (
	"log"

	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/model"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/repository"
)

type personaService struct {
	personaRepo repository.PersonaRepository
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
func NewPersonaService(pR repository.PersonaRepository) PersonaService {
	return &personaService{pR}
}

func (pB personaService) CreaPersona(nom string, ape string, fechaNacimiento string) model.Personas {

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

func (pB personaService) BuscaTodo() model.Personas {

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

func (pB personaService) BorraPersona(id string) model.Personas {

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

func (pB personaService) BuscaPersona(id string) model.Personas {

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

func (pB personaService) ActualizaPersona(id string, nom string, ape string, fechaNacimiento string) model.Personas {

	p := &model.Persona{Nombre: nom, Apellido: ape, FechaNacimiento: fechaNacimiento, ID: id}
	p, err := pB.personaRepo.Actualiza(p)
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