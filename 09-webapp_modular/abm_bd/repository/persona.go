package repository

import (
	"database/sql"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/model"
	"log"
	"strconv"
	"time"
)

//TODO agregar los null
//y unique a Usuario en BD
type Persona struct {
	ID              int
	Nombre          string
	Apellido        string
	FechaNacimiento sql.NullTime
}

type PersonaRepository struct {
	db *sql.DB
}

func NewPersonaRepository(db *sql.DB) *PersonaRepository {
	return &PersonaRepository{db}
}

func (pR PersonaRepository) Persiste(p model.Persona) (model.Persona, error) {
	//Inicio como nulo, si no lo es lo cambio
	fechaNull := sql.NullTime{
		Valid: false,
	}
	if p.FechaNacimiento != "" {
		fecha, err := time.Parse("02-01-2006", p.FechaNacimiento)
		if err != nil {
			log.Println("Error persiste Persona con fecha:", err)
			return model.Persona{}, err
		}
		fechaNull = sql.NullTime{
			Time:  fecha,
			Valid: true,
		}
	}
	stmt, err := pR.db.Prepare("INSERT into persona(nombre, apellido, fecha_nacimiento) VALUES(?,?,?);")
	if err != nil {
		return model.Persona{}, err
	}

	_, err = stmt.Exec(p.Nombre, p.Apellido, fechaNull)
	if err != nil {
		return model.Persona{}, err
	}

	return p, nil
}

func (pR PersonaRepository) Borra(id string) error {

	stmt, err := pR.db.Prepare("DELETE FROM persona WHERE id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (pR PersonaRepository) BuscaTodo() ([]model.Persona, error) {

	rows, err := pR.db.Query("SELECT id, nombre, apellido, fecha_nacimiento FROM persona;")
	if err != nil {
		return []model.Persona{}, err
	}
	defer rows.Close()

	var rP []model.Persona
	var p Persona

	var fecha string
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Nombre, &p.Apellido, &p.FechaNacimiento)
		if err != nil {
			return []model.Persona{}, err
		}

		if p.FechaNacimiento.Valid == true {
			fecha = p.FechaNacimiento.Time.Format("02-01-2006")
		} else {
			fecha = ""
		}

		per := model.Persona{ID: strconv.Itoa(p.ID),
			Nombre:          p.Nombre,
			Apellido:        p.Apellido,
			FechaNacimiento: fecha,
		}

		rP = append(rP, per)
	}

	return rP, nil
}

func (pR PersonaRepository) BuscaPorId(id string) (model.Persona, error) {
	var p Persona
	err := pR.db.QueryRow("SELECT id, nombre, apellido, fecha_nacimiento FROM persona WHERE id = ?;", id).
		Scan(&p.ID, &p.Nombre, &p.Apellido, &p.FechaNacimiento)
	if err != nil {
		return model.Persona{}, err
	}

	var fecha string
	if p.FechaNacimiento.Valid == true {
		fecha = p.FechaNacimiento.Time.Format("02-01-2006")
	} else {
		fecha = ""
	}
	perM := model.Persona{ID: strconv.Itoa(p.ID),
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: fecha,
	}
	return perM, nil
}

func (pR PersonaRepository) Actualiza(p model.Persona) (model.Persona, error) {

	//Inicio como nulo, si no lo es lo cambio
	fechaNull := sql.NullTime{
		Valid: false,
	}
	if p.FechaNacimiento != "" {
		fecha, err := time.Parse("02-01-2006", p.FechaNacimiento)
		if err != nil {
			log.Println("Error actualiza Persona con fecha:", err)
			return model.Persona{}, err
		}
		fechaNull = sql.NullTime{
			Time:  fecha,
			Valid: true,
		}
	}

	stmt, err := pR.db.Prepare("Update persona SET nombre=?, apellido=?, fecha_nacimiento=? WHERE id=?;")
	if err != nil {
		return model.Persona{}, err
	}

	_, err = stmt.Exec(p.Nombre, p.Apellido, fechaNull, p.ID)
	if err != nil {
		return model.Persona{}, err
	}

	return p, nil
}
