package repository

import (
	"database/sql"
	"log"
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

func (pR PersonaRepository) Persiste(p Persona) (Persona, error) {
	stmt, err := pR.db.Prepare("INSERT into persona(nombre, apellido, fecha_nacimiento) VALUES(?,?,?);")
	if err != nil {
		return Persona{}, err
	}

	_, err = stmt.Exec(p.Nombre, p.Apellido, p.FechaNacimiento)
	if err != nil {
		return Persona{}, err
	}

	return p, nil
}

func (pR PersonaRepository) Borra(id int) error {

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

func (pR PersonaRepository) BuscaTodo() ([]Persona, error) {

	rows, err := pR.db.Query("SELECT id, nombre, apellido, fecha_nacimiento FROM persona;")
	if err != nil {
		return []Persona{}, err
	}
	defer rows.Close()

	var rP []Persona
	var p Persona
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Nombre, &p.Apellido, &p.FechaNacimiento)
		if err != nil {
			return []Persona{}, err
		}
		rP = append(rP, p)
	}
	return rP, nil
}

func (pR PersonaRepository) BuscaPorId(id int) (Persona, error) {
	var p Persona
	err := pR.db.QueryRow("SELECT id, nombre, apellido, fecha_nacimiento FROM persona WHERE id = ?;", id).
		Scan(&p.ID, &p.Nombre, &p.Apellido, &p.FechaNacimiento)
	if err != nil {
		return Persona{}, err
	}
	return p, nil
}

func (pR PersonaRepository) Actualiza(p Persona) (Persona, error) {
	stmt, err := pR.db.Prepare("Update persona SET nombre=?, apellido=?, fecha_nacimiento=? WHERE id=?;")
	if err != nil {
		return Persona{}, err
	}

	borrar, err := stmt.Exec(p.Nombre, p.Apellido, p.FechaNacimiento, p.ID)
	if err != nil {
		return Persona{}, err
	}
	log.Println("Error actualiza Persona con fecha:", p, err, borrar)
	rowCnt, err := borrar.RowsAffected()
	log.Println("Error actualiza Persona con fechazzzz:", rowCnt, err)
	return p, nil
}
