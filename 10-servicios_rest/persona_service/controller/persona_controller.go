package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mberliner/gobase/10-servicios_rest/persona_service/domain"
	"github.com/mberliner/gobase/10-servicios_rest/persona_service/logger"
	"github.com/mberliner/gobase/10-servicios_rest/persona_service/service"
)

// PersonaController interface del controller de Persona
type PersonaController interface {
	BuscarTodo(c *gin.Context)
	Crear(c *gin.Context)
	BuscarPorId(c *gin.Context)
	Borrar(c *gin.Context)
	Actualizar(c *gin.Context)
	ActualizarParcial(c *gin.Context)
}

type personaController struct {
	personaService service.PersonaService
}

// NewPersonaController para acceder a controller de persona de forma ordenada
func NewPersonaController(pS service.PersonaService) PersonaController {
	return &personaController{pS}
}

func (pC personaController) BuscarTodo(c *gin.Context) {

	personas, err := pC.personaService.BuscaTodo()

	if err != nil {
		sError := NewInternalServerError("Error al buscar Personas", err)
		c.JSON(sError.Status(), sError)
		return
	}
	c.JSON(http.StatusOK, personas)

}

func (pC personaController) Crear(c *gin.Context) {
	var persona domain.Persona
	if err := c.ShouldBindJSON(&persona); err != nil {
		logger.Error("Error en Crear persona con parametros:", err, persona)
		restErr := NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	p, err := pC.personaService.CreaPersona(&persona)
	if err != nil {
		sError := NewInternalServerError("Error al crear Persona", err)
		c.JSON(sError.Status(), sError)
		return
	}

	c.JSON(http.StatusCreated, p)

}

func (pC personaController) BuscarPorId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Error al convertir ID", err, id)
		sError := NewBadRequestError("Error al convertir ID")
		c.JSON(sError.Status(), sError)
		return
	}

	p, err := pC.personaService.BuscaPersona(id)
	if err != nil {
		if strings.Contains(err.Error(), "Not Found") {
			sError := NewNotFoundError("Error al buscar Persona por ID, Not Found")
			c.JSON(sError.Status(), sError)

		} else {
			sError := NewInternalServerError("Error al buscar Persona por ID", err)
			c.JSON(sError.Status(), sError)
		}
		return
	}

	c.JSON(http.StatusCreated, p)

}

func (pC personaController) Borrar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Error al convertir ID", err, id)
		sError := NewBadRequestError("Error al convertir ID, Borrar")
		c.JSON(sError.Status(), sError)
		return
	}
	if err := pC.personaService.Borra(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func (pC personaController) Actualizar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Error al convertir ID, Actualizar", err, id)
		sError := NewBadRequestError("Error al convertir ID")
		c.JSON(sError.Status(), sError)
		return
	}

	var persona domain.Persona

	if err := c.ShouldBindJSON(&persona); err != nil {
		logger.Error("Error en Actualizar persona, parametros", err, persona)
		restErr := NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	persona.ID = id
	p, err := pC.personaService.Actualiza(&persona)
	if err != nil {
		if strings.Contains(err.Error(), "Not Found") {
			sError := NewNotFoundError("Error al Actualizar Persona por ID, Not Found")
			c.JSON(sError.Status(), sError)

		} else {
			sError := NewInternalServerError("Error al Actualizar Persona por ID", err)
			c.JSON(sError.Status(), sError)
		}
		return
	}

	c.JSON(http.StatusOK, p)

}

func (pC personaController) ActualizarParcial(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Error en Actualizar Parcial persona, id", err, id)
		sError := NewBadRequestError("Error ID")
		c.JSON(sError.Status(), sError)
		return
	}

	var persona domain.Persona

	if err := c.ShouldBindJSON(&persona); err != nil {
		logger.Error("Error en Actualizar Parcial persona, parametros", err, persona)
		restErr := NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	persona.ID = id
	p, err := pC.personaService.ActualizaParcial(&persona)
	if err != nil {
		if strings.Contains(err.Error(), "Not Found") {
			sError := NewNotFoundError("Error al Actualizar Parcial Persona por ID, Not Found")
			c.JSON(sError.Status(), sError)

		} else {
			sError := NewInternalServerError("Error al Actualizar Parcial Persona por ID", err)
			c.JSON(sError.Status(), sError)
		}
		return
	}

	c.JSON(http.StatusOK, p)

}
