package business

import (
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/model"
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/repository"
	"log"
)

func CreaPersona(nom string, ape string, fechaNacimiento string) model.Personas {

	p := &model.Persona{Nombre: nom, Apellido: ape, FechaNacimiento: fechaNacimiento}
	p, err := repository.PR.Persiste(p)
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

func BuscaTodo() model.Personas {

	ps, err := repository.PR.BuscaTodo()
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

func BorraPersona(id string) model.Personas {

	err := repository.PR.Borra(id)
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

func BuscaPersona(id string) model.Personas {

	p, err := repository.PR.BuscaPorID(id)
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

func ActualizaPersona(id string, nom string, ape string, fechaNacimiento string) model.Personas {

	p := &model.Persona{Nombre: nom, Apellido: ape, FechaNacimiento: fechaNacimiento, ID: id}
	p, err := repository.PR.Actualiza(p)
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
