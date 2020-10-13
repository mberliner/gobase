package business

import (
	//	"errors"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/model"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/repository"
	"log"
)

func CreaPersona(nom string, ape string, fechaNacimiento string) model.Personas {

	p := repository.Persona{Nombre: nom, Apellido: ape}
	p, err := repository.PR.Persiste(p)
	if err != nil {
		log.Println("Error persiste Pesona:", err)
		mP := model.Personas{}
		mP.Error = err
		return mP
	}
	per := model.Persona{ID: p.ID,
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: p.FechaNacimiento,
	}

	personas := []model.Persona{per}

	mP := model.Personas{
		PersonasM: personas,
		Error:     nil,
		Mensaje:   "Persona Creada ok",
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
	personas := []model.Persona{}
	for _, p := range ps {

		per := model.Persona{ID: p.ID,
			Nombre:          p.Nombre,
			Apellido:        p.Apellido,
			FechaNacimiento: p.FechaNacimiento,
		}
		personas = append(personas, per)
	}
	mP := model.Personas{
		PersonasM: personas,
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
