package repository

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/mberliner/gobase/10-servicios_rest/entities_service/domain"
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

type UserRepository interface {
	Persiste(u *domain.User) (*domain.User, error)
	BuscaPorUsuario(usu string) (*domain.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (uR userRepository) Persiste(u *domain.User) (*domain.User, error) {

	edadNull := sql.NullInt64{
		Valid: false,
	}
	if u.Edad != "" {
		edadI, err := strconv.Atoi(u.Edad)
		if err != nil {
			log.Println("Error edad debe ser numerico:", err)
			return &domain.User{}, err
		}
		edadNull = sql.NullInt64{
			Int64: int64(edadI),
			Valid: true,
		}
	}

	stmt, err := uR.db.Prepare("INSERT into user(usuario, nombre, apellido, edad, password) VALUES(?,?,?,?, ?);")
	if err != nil {
		return &domain.User{}, err
	}

	_, err = stmt.Exec(u.Usuario, u.Nombre, u.Apellido, edadNull, u.Password)
	if err != nil {
		return &domain.User{}, err
	}

	return u, nil
}

func (uR userRepository) BuscaPorUsuario(usu string) (*domain.User, error) {
	//usuario es unique
	rows, err := uR.db.Query("SELECT id, usuario, edad, nombre, apellido, password  FROM user WHERE usuario = ?;", usu)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var u user
	var edad string
	var uD *domain.User
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Usuario, &u.Edad, &u.Nombre, &u.Apellido, &u.Password)
		if err != nil {
			return nil, err
		}

		if u.Edad.Valid == true {
			edad = strconv.Itoa(int(u.Edad.Int64))
		} else {
			edad = ""
		}

		uD = &domain.User{
			ID:       u.ID,
			Usuario:  u.Usuario,
			Nombre:   u.Nombre,
			Apellido: u.Apellido,
			Edad:     edad,
			Password: u.Password}
	}
	return uD, nil
}
