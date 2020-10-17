package repository

import (
	"database/sql"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/model"
)

//TODO agregar los null
//y unique a Usuario en BD
type user struct {
	ID       int
	Usuario  string
	Nombre   string
	Apellido string
	Edad     sql.NullInt64
	Password string
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (uR UserRepository) Persiste(u model.User) (model.User, error) {
	stmt, err := uR.db.Prepare("INSERT into user(usuario, nombre, apellido, edad, password) VALUES(?,?,?,?);")
	if err != nil {
		return model.User{}, err
	}

	_, err = stmt.Exec(u.Usuario, u.Nombre, u.Apellido, u.Password, u.Password)
	if err != nil {
		return model.User{}, err
	}

	return u, nil
}

func (uR UserRepository) BuscaPorUsuario(usu string) ([]model.User, error) {
	rows, err := uR.db.Query("SELECT id, usuario, edad, nombre, apellido, password  FROM user WHERE usuario = ?;", usu)
	if err != nil {
		return []model.User{}, err
	}
	defer rows.Close()

	var rU []model.User
	var u user
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Usuario, &u.Edad, &u.Nombre, &u.Apellido, &u.Password)
		if err != nil {
			return []model.User{}, err
		}
		uM := model.User{
			ID:       u.ID,
			Usuario:  u.Usuario,
			Nombre:   u.Nombre,
			Apellido: u.Apellido,
			Edad:     u.Edad,
			Password: u.Password,
			Error:    nil}
		rU = append(rU, uM)
	}
	return rU, nil
}
