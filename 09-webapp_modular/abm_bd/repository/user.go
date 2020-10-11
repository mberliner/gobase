package repository

import (
	"database/sql"
)

//TODO agregar los null
//y unique a Usuario en BD
type User struct {
	ID       int
	Usuario  string
	Nombre   string
	Apellido string
	Edad     sql.NullInt64
	Password []byte
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (uR UserRepository) Persiste(u User) (User, error) {
	stmt, err := uR.db.Prepare("INSERT into user(usuario, nombre, apellido, password) VALUES(?,?,?,?);")
	if err != nil {
		return User{}, err
	}

	_, err = stmt.Exec(u.Usuario, u.Nombre, u.Apellido, string(u.Password))
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (uR UserRepository) BuscaPorUsuario(usu string) ([]User, error) {
	var u User
	rows, err := uR.db.Query("SELECT id, edad, nombre, apellido, password  FROM user WHERE usuario = ?;", usu)
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	var pass string

	var rU []User
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Edad, &u.Nombre, &u.Apellido, &pass)
		if err != nil {
			return []User{}, err
		}
		//TODO Revisar
		u.Password = []byte(pass)
		u.Usuario = usu
		rU = append(rU, u)
	}
	return rU, nil
}
