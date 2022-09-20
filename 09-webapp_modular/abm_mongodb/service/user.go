package service

import (
	"errors"
	"log"

	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/model"
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/repository/mongoDB"
	"golang.org/x/crypto/bcrypt"
)

// UserService interface para exponer manejo de User
type UserService interface {
	CreaUsuario(usu string, pass string, nom string, ape string) model.User
	Autentica(usu string, pass string) (model.User, bool)
	BuscaPorUsuario(usu string) model.User
}

// NewUserService para obtener repositorio de manera ordenada
func NewUserService(uR mongoDB.UserRepository) UserService {
	return &userService{
		userRepo: uR,
	}
}

type userService struct {
	userRepo mongoDB.UserRepository
}

func (userS userService) CreaUsuario(usu string, pass string, nom string, ape string) model.User {

	sU, err := userS.userRepo.BuscaPorUsuario(usu)
	if len(sU) > 0 {
		mU := model.User{}
		mU.Error = errors.New("el usuario ya existe, elija otro nombre ")
		return mU
	}
	if err != nil {
		mU := model.User{}
		//TODO poner error de negocio
		mU.Error = err
		log.Println("Error al buscar usuario", mU, err, len(sU))
		return mU
	}
	encrPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		mU := model.User{}
		mU.Error = err
		return mU
	}

	u := &model.User{Usuario: usu, Nombre: nom, Apellido: ape, Password: string(encrPass)}
	u, err = userS.userRepo.Persiste(u)
	if err != nil {
		log.Println("Error persiste:", err)
		mU := model.User{}
		mU.Error = err
		return mU
	}
	mU := model.User{
		ID:       u.ID,
		Usuario:  u.Usuario,
		Nombre:   u.Nombre,
		Apellido: u.Apellido,
		Edad:     u.Edad,
		Password: u.Password,
		Mensaje:  "Usuario creado con éxito, ya puede ir al login", //al controller?
		Error:    nil}
	return mU
}

func (userS userService) Autentica(usu string, pass string) (model.User, bool) {
	sU, err := userS.userRepo.BuscaPorUsuario(usu)
	if err != nil {
		log.Println("BuscaporUsuario:", sU, err)
		mU := model.User{}
		mU.Error = err
		return mU, false
	}
	if len(sU) == 0 {
		mU := model.User{}
		mU.Error = errors.New("el usuario o password es incorrecto ")
		return mU, false
	}

	err = bcrypt.CompareHashAndPassword([]byte(sU[0].Password), []byte(pass))
	if err != nil {
		mU := model.User{}
		mU.Error = errors.New("el usuario o password es incorrecto ")
		return mU, false

	}
	mU := model.User{
		ID:       sU[0].ID,
		Usuario:  sU[0].Usuario,
		Nombre:   sU[0].Nombre,
		Apellido: sU[0].Apellido,
		Edad:     sU[0].Edad,
		Password: sU[0].Password,
		Mensaje:  "Autenticación exitosa",
		Error:    nil}

	return mU, true
}

func (userS userService) BuscaPorUsuario(usu string) model.User {

	sU, err := userS.userRepo.BuscaPorUsuario(usu)
	if err != nil || len(sU) == 0 {
		mU := model.User{}
		mU.Error = errors.New("usuario no encontrado")
		return mU
		//TODO log y revisar
	}
	mU := model.User{
		ID:       sU[0].ID,
		Usuario:  sU[0].Usuario,
		Nombre:   sU[0].Nombre,
		Apellido: sU[0].Apellido,
		Edad:     sU[0].Edad,
		Password: sU[0].Password,
		Mensaje:  "Usuario OK",
		Error:    nil}

	return mU
}
