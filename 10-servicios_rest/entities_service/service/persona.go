package service

import (
	"log"

	"github.com/mberliner/gobase/10-servicios_rest/entities_service/domain"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/repository"
)

type personaService struct {
	personaRepo repository.PersonaRepository
}

//PersonaService interface para poder realizar tests del negocio de persona
type PersonaService interface {
	CreaPersona(*domain.Persona) (*domain.Persona, error)
	BuscaTodo() ([]domain.Persona, error)
	Borra(id int) error
	BuscaPersona(id int) (*domain.Persona, error)
	Actualiza(*domain.Persona) (*domain.Persona, error)
}

//NewPersonaService para obtener megocio de forma ordenada
func NewPersonaService(pR repository.PersonaRepository) PersonaService {
	return &personaService{pR}
}

func (pB personaService) CreaPersona(per *domain.Persona) (*domain.Persona, error) {

	p, err := pB.personaRepo.Persiste(per)
	if err != nil {
		log.Println("Error persiste Pesona:", err)
		return nil, err
	}

	return p, nil
}

func (pS personaService) BuscaTodo() ([]domain.Persona, error) {

	ps, err := pS.personaRepo.BuscaTodo()
	if err != nil {
		log.Println("Error buscaTodo:", err)
		return nil, err
	}

	return ps, nil
}

func (pS personaService) Borra(id int) error {

	err := pS.personaRepo.Borra(id)
	if err != nil {
		log.Println("Error borraPersona:", err, id)
		return err
	}

	return nil
}

func (pB personaService) BuscaPersona(id int) (*domain.Persona, error) {

	p, err := pB.personaRepo.BuscaPorID(id)
	if err != nil {
		log.Println("Error BuscaPersona:", err, "id:", id)
		return nil, err
	}

	return p, nil
}

func (pB personaService) Actualiza(per *domain.Persona) (*domain.Persona, error) {

	p, err := pB.personaRepo.Actualiza(per)
	if err != nil {
		log.Println("Error actualiza Persona:", err)
		return nil, err
	}

	return p, nil
}
