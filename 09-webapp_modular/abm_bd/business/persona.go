package business

import (
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/model"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/repository"
	"log"
	"strconv"
)

func CreaPersona(nom string, ape string, fechaNacimiento string) model.Personas {

	p := model.Persona{Nombre: nom, Apellido: ape, FechaNacimiento: fechaNacimiento}
	p, err := repository.PR.Persiste(p)
	if err != nil {
		log.Println("Error persiste Pesona:", err)
		mP := model.Personas{}
		mP.Error = err
		return mP
	}

	personas := []model.Persona{p}

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

func BorraPersona(id int) model.Personas {

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

func BuscaPersona(id int) model.Personas {

	p, err := repository.PR.BuscaPorId(id)
	log.Println("Error en editarPersona1-------:", id, p, err)
	if err != nil {
		log.Println("Error buscaPersona:", err)
		mP := model.Personas{}
		mP.Error = err
		return mP
	}

	mP := model.Personas{
		PersonasM: []model.Persona{p},
		Error:     nil,
		Mensaje:   "Busca ok",
	}
	log.Println("Error en editarPersona2-------:", mP)
	return mP
}

func ActualizaPersona(id string, nom string, ape string, fechaNacimiento string) model.Personas {

	idd, _ := strconv.Atoi(id)
	p := model.Persona{Nombre: nom, Apellido: ape, FechaNacimiento: fechaNacimiento, ID: idd}
	p, err := repository.PR.Actualiza(p)
	log.Println("Error actualiza Persona con fecha:", p, err)
	if err != nil {
		log.Println("Error actualiza Pesona:", err)
		mP := model.Personas{}
		mP.Error = err
		return mP
	}

	mP := model.Personas{
		PersonasM: []model.Persona{p},
		Error:     nil,
		Mensaje:   "Persona Actualizada Ok",
	}
	return mP
}
