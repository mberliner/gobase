package mongoDBOld

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/mberliner/gobase/09-webapp_modular/abm/model"
)

type user struct {
	ID       bson.ObjectId //`json:"id" bson:"id"`
	Usuario  string        //`json:"Usuario" bson:"Usuario"`
	Nombre   string        //`json:"Nombre" bson:"Nombre"`
	Apellido string        //`json:"Apellido" bson:"Apellido"`
	Edad     string        //`json:"Edad" bson:"Edad"`
	Password string        //`json:"Password" bson:"Password"`
}

// UserRepository interface para acceso a datos de User
type UserRepository interface {
	Persiste(u *model.User) (*model.User, error)
	BuscaPorUsuario(usu string) ([]model.User, error)
}

// NewUserRepository para obtener repositorio de manera ordenada
func NewUserRepository(db *mgo.Database) UserRepository {
	return &userRepository{db}
}

type userRepository struct {
	db *mgo.Database
}

func (uR userRepository) Persiste(u *model.User) (*model.User, error) {

	uL := user{ID: bson.NewObjectId(),
		Usuario:  u.Usuario,
		Nombre:   u.Nombre,
		Apellido: u.Apellido,
		Edad:     u.Edad,
		Password: u.Password,
	}
	err := uR.db.C("user").Insert(uL)
	if err != nil {
		return &model.User{}, err
	}

	return u, nil
}

func (uR userRepository) BuscaPorUsuario(usu string) ([]model.User, error) {

	u := user{}
	err := uR.db.C("user").Find(bson.M{"usuario": usu}).One(&u)
	if err != nil && err.Error() == "not found" {
		return nil, nil
	}
	if err != nil {
		return []model.User{}, err
	}

	uM := model.User{
		ID:       u.ID.String(),
		Usuario:  u.Usuario,
		Nombre:   u.Nombre,
		Apellido: u.Apellido,
		Edad:     u.Edad,
		Password: u.Password,
		Error:    nil}
	suM := []model.User{}
	suM = append(suM, uM)

	return suM, nil
}
