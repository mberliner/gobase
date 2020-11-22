package service

import (
	"errors"

	"github.com/mberliner/gobase/10-servicios_rest/entities_service/domain"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/logger"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/repository"
	"golang.org/x/crypto/bcrypt"
)

//UserService interface para exponer manejo de User
type UserService interface {
	CreaUsuario(*domain.User) (*domain.User, error)
	Autentica(usu string, pass string) (*domain.User, error)
	BuscaPorUsuario(usu string) (*domain.User, error)
}

//NewUserService para obtener repositorio de manera ordenada
func NewUserService(uR repository.UserRepository) UserService {
	return &userService{
		userRepo: uR,
	}
}

type userService struct {
	userRepo repository.UserRepository
}

func (uS userService) CreaUsuario(user *domain.User) (*domain.User, error) {

	sU, err := uS.userRepo.BuscaPorUsuario(user.Usuario)
	if sU != nil {
		eU := errors.New("El usuario ya existe")
		return nil, eU
	}
	if err != nil {
		return nil, err
	}
	encrPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		logger.Error("Error encriptar", err)
		return nil, err
	}

	u := &domain.User{Usuario: user.Usuario,
		Nombre:   user.Nombre,
		Apellido: user.Apellido,
		Password: string(encrPass),
	}
	mU, err := repository.UserRepo.Persiste(u)
	if err != nil {
		return nil, err
	}
	return mU, nil
}

func (uS userService) Autentica(usu string, pass string) (*domain.User, error) {
	sU, err := uS.userRepo.BuscaPorUsuario(usu)
	if err != nil {
		logger.Error("Error BuscaporUsuario: "+usu, err)
		return nil, err
	}
	if sU == nil {
		logger.Error("Error BuscaporUsuario: No existe Usuario: "+usu, nil)
		return nil, errors.New("El usuario o password es incorrecto ")
	}

	err = bcrypt.CompareHashAndPassword([]byte(sU.Password), []byte(pass))
	if err != nil {
		logger.Error("Error BuscaporUsuario: Pass incorrecta:", err)
		return nil, errors.New("El usuario o password es incorrecto ")
	}

	return sU, nil
}

func (uS userService) BuscaPorUsuario(usu string) (*domain.User, error) {

	sU, err := uS.userRepo.BuscaPorUsuario(usu)
	if err != nil {
		return nil, err
	}
	if sU == nil {

		return nil, errors.New("Usuario no encontrado")
		//TODO log y revisar
	}

	return sU, nil
}
