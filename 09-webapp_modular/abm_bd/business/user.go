package business

import (
	"errors"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/model"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreaUsuario(usu string, pass string, nom string, ape string) model.User {

	sU, err := repository.UR.BuscaPorUsuario(usu)
	if err != nil || len(sU) > 0 {
		mU := model.User{}
		mU.Error = errors.New("El usuario ya existe, elija otro nombre ")
		log.Println("El usuario ya existe, elija otro nombre:---------", mU, err, len(sU))
		return mU
	}
	encrPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		mU := model.User{}
		mU.Error = err
		return mU
	}

	u := repository.User{Usuario: usu, Nombre: nom, Apellido: ape, Password: encrPass}
	u, err = repository.UR.Persiste(u)
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
		Password: string(u.Password),
		Mensaje:  "Usuario creado con éxito, ya puede ir al login",
		Error:    nil}
	return mU
}

func Autentica(usu string, pass string) (model.User, bool) {
	sU, err := repository.UR.BuscaPorUsuario(usu)
	if err != nil || len(sU) == 0 {
		log.Println("BuscaporUsuario:", sU, err)
		mU := model.User{}
		mU.Error = errors.New("El usuario o password es incorrecto ")
		return mU, false
	}

	err = bcrypt.CompareHashAndPassword(sU[0].Password, []byte(pass))
	if err != nil {
		mU := model.User{}
		mU.Error = errors.New("El usuario o password es incorrecto ")
		return mU, false

	}
	mU := model.User{
		ID:       sU[0].ID,
		Usuario:  sU[0].Usuario,
		Nombre:   sU[0].Nombre,
		Apellido: sU[0].Apellido,
		Edad:     sU[0].Edad,
		Password: string(sU[0].Password),
		Mensaje:  "Autenticación exitosa",
		Error:    nil}

	return mU, true
}

func BuscaPorUsuario(usu string) model.User {

	sU, err := repository.UR.BuscaPorUsuario(usu)
	if err != nil || len(sU) == 0 {
		mU := model.User{}
		mU.Error = errors.New("Usuario no encontrado")
		return mU
		//TODO log y revisar
	}
	mU := model.User{
		ID:       sU[0].ID,
		Usuario:  sU[0].Usuario,
		Nombre:   sU[0].Nombre,
		Apellido: sU[0].Apellido,
		Edad:     sU[0].Edad,
		Password: string(sU[0].Password),
		Mensaje:  "Usuario OK",
		Error:    nil}

	return mU
}
