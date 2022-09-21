package mysql

import (
	"database/sql"
	"strconv"

	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/model"
)

// TODO agregar los null
// y unique a Usuario en BD
type user struct {
	ID       string
	Usuario  string
	Nombre   string
	Apellido string
	Edad     sql.NullInt64
	Password string
}

type UserRepository interface {
	Persiste(u *model.User) (*model.User, error)
	BuscaPorUsuario(usu string) ([]model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (uR userRepository) Persiste(u *model.User) (*model.User, error) {

	edadNull := sql.NullInt64{
		Valid: false,
	}
	if u.Edad != "" {
		edadI, err := strconv.Atoi(u.Edad)
		if err != nil {
			//log.Println("Error edad debe ser numerico:", err)
			return &model.User{}, err
		}
		edadNull = sql.NullInt64{
			Int64: int64(edadI),
			Valid: true,
		}
	}

	stmt, err := uR.db.Prepare("INSERT into user(usuario, nombre, apellido, edad, password) VALUES(?,?,?,?, ?);")
	if err != nil {
		return &model.User{}, err
	}

	_, err = stmt.Exec(u.Usuario, u.Nombre, u.Apellido, edadNull, u.Password)
	if err != nil {
		return &model.User{}, err
	}

	return u, nil
}

func (uR userRepository) BuscaPorUsuario(usu string) ([]model.User, error) {
	rows, err := uR.db.Query("SELECT id, usuario, edad, nombre, apellido, password  FROM user WHERE usuario = ?;", usu)
	if err != nil {
		return []model.User{}, err
	}
	defer rows.Close()

	var rU []model.User
	var u user
	var edad string
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Usuario, &u.Edad, &u.Nombre, &u.Apellido, &u.Password)
		if err != nil {
			return []model.User{}, err
		}

		if u.Edad.Valid == true {
			edad = strconv.Itoa(int(u.Edad.Int64))
		} else {
			edad = ""
		}

		uM := model.User{
			ID:       u.ID,
			Usuario:  u.Usuario,
			Nombre:   u.Nombre,
			Apellido: u.Apellido,
			Edad:     edad,
			Password: u.Password,
			Error:    nil}
		rU = append(rU, uM)
	}
	return rU, nil
}
