package business

import (
	"errors"
	"fmt"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/repository"
	"golang.org/x/crypto/bcrypt"
)

func CreaUsuario(usu string, pass string, nom string, ape string) (repository.User, error) {

	sU, err := repository.UR.BuscaPorUsuario(usu)
	if err != nil || len(sU) > 0 {
		return repository.User{}, errors.New("El usuario ya existe, elija otro nombre ")
	}
	encrPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return repository.User{}, err
	}

	u := repository.User{Usuario: usu, Nombre: nom, Apellido: ape, Password: encrPass}
	_, err = repository.UR.Persiste(u)
	if err != nil {
		//TODO log
		fmt.Println("err persiste:", err)
		return repository.User{}, err
	}
	return u, nil
}

func Autentica(usu string, pass string) bool {
	sU, err := repository.UR.BuscaPorUsuario(usu)
	if err != nil || len(sU) == 0 {
		fmt.Println("user:------------------entra2-------------------------", sU, err)
		return false
	}

	err = bcrypt.CompareHashAndPassword(sU[0].Password, []byte(pass))
	if err != nil {
		fmt.Println("user:------------------entra3-------------------------", sU, err, pass, sU[0].Password)
		return false
	}

	return true
}
