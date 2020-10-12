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
	stmt, err := pR.db.Prepare("INSERT into persona(nombre, apellido, fechanacimiento) VALUES(?,?,?);")
	if err != nil {
		return Persona{}, err
	}

	_, err = stmt.Exec(p.Nombre, p.Apellido, nil)
	if err != nil {
		return Persona{}, err
	}

	return p, nil
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

func (pR PersonaRepository) buscaPorID(id int) (Persona, error) {
	var p Persona
	err := pR.db.QueryRow("SELECT id, nombre, apellido, fechanacimiento FROM persona WHERE id = ?;", id).
		Scan(&p.ID, &p.FechaNacimiento, &p.Nombre, &p.Apellido)
	if err != sql.ErrNoRows {
		return Persona{}, err
	}

	return p, nil
}

/*
func (pR PersonaRepository) BuscaPorUsuario(usu string) ([]Persona, error) {
	var u Persona
	rows, err := pR.db.Query("SELECT id, edad, nombre, apellido, password  FROM user WHERE usuario = ?;", usu)
	if err != nil {
		return []Persona{}, err
	}
	defer rows.Close()

	var rU []Persona
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Edad, &p.Nombre, &p.Apellido, &pass)
		if err != nil {
			return []Persona{}, err
		}
		//TODO Revisar
		p.Password = []byte(pass)
		p.Usuario = usu
		rU = append(rU, u)
	}
	return rU, nil
}
*/
