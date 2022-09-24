package service

import (
	"github.com/mberliner/gobase/10-servicios_rest/persona_service/domain"
	"github.com/mberliner/gobase/10-servicios_rest/persona_service/logger"
	"github.com/mberliner/gobase/10-servicios_rest/persona_service/repository"
)

type personaService struct {
	personaRepo repository.PersonaRepository
}

// PersonaService interface para poder realizar tests del negocio de persona
type PersonaService interface {
	CreaPersona(*domain.Persona) (*domain.Persona, error)
	BuscaTodo() ([]domain.Persona, error)
	Borra(int) error
	BuscaPersona(id int) (*domain.Persona, error)
	Actualiza(*domain.Persona) (*domain.Persona, error)
	ActualizaParcial(*domain.Persona) (*domain.Persona, error)
}

// NewPersonaService para obtener megocio de forma ordenada
func NewPersonaService(pR repository.PersonaRepository) PersonaService {
	return &personaService{pR}
}

func (pS personaService) CreaPersona(per *domain.Persona) (*domain.Persona, error) {

	p, err := pS.personaRepo.Persiste(per)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (pS personaService) BuscaTodo() ([]domain.Persona, error) {

	ps, err := pS.personaRepo.BuscaTodo()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (pS personaService) Borra(id int) error {

	err := pS.personaRepo.Borra(id)
	if err != nil {
		return err
	}

	return nil
}

func (pS personaService) BuscaPersona(id int) (*domain.Persona, error) {

	p, err := pS.personaRepo.BuscaPorID(id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (pS personaService) Actualiza(per *domain.Persona) (*domain.Persona, error) {

	p, err := pS.personaRepo.Actualiza(per)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (pS personaService) ActualizaParcial(per *domain.Persona) (*domain.Persona, error) {

	pBuscada, err := pS.personaRepo.BuscaPorID(per.ID)
	if err != nil {
		logger.Error("Error actualizar parcial Persona (BuscaPorId):", err)
		return nil, err
	}
	if per.Apellido != "" {
		pBuscada.Apellido = per.Apellido
	}
	if per.Nombre != "" {
		pBuscada.Nombre = per.Nombre
	}
	if per.FechaNacimiento != "" {
		pBuscada.FechaNacimiento = per.FechaNacimiento
	}

	p, err := pS.personaRepo.Actualiza(pBuscada)
	if err != nil {
		logger.Error("Error actualizar parcial Persona (Actualiza):", err)
		return nil, err
	}

	return p, nil
}
