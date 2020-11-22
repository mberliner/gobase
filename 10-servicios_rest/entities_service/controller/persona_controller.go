package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/domain"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/logger"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/service"
)

//PersonaController interface del controller de Persona
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

//NewPersonaController para acceder a controller de persona de forma ordenada
func NewPersonaController(pS service.PersonaService) PersonaController {
	return &personaController{pS}
}

func (pC personaController) BuscarTodo(c *gin.Context) {

	personas, err := pC.personaService.BuscaTodo()

	if err != nil {
		sError := NewInternalServerError("Error al buscar Personas", err)
		c.JSON(http.StatusInternalServerError, sError)
		return
	}
	c.JSON(http.StatusOK, personas)

}

func (pC personaController) Crear(c *gin.Context) {
	var persona domain.Persona

	if err := c.ShouldBindJSON(&persona); err != nil {
		logger.Error("Error en Crear persona con parametros:", err)
		restErr := NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	p, err := pC.personaService.CreaPersona(&persona)
	if err != nil {
		sError := NewInternalServerError("Error al crear Persona", err)
		c.JSON(http.StatusInternalServerError, sError)
		return
	}

	c.JSON(http.StatusCreated, p)

}

func (pC personaController) BuscarPorId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Error al convertir ID", err)
		sError := NewBadRequestError("Error al convertir ID")
		c.JSON(http.StatusInternalServerError, sError)
		return
	}

	p, err := pC.personaService.BuscaPersona(id)
	if err != nil {
		sError := NewInternalServerError("Error al buscar Persona por ID", err)
		c.JSON(http.StatusInternalServerError, sError)
		return
	}

	c.JSON(http.StatusCreated, p)

}

func (pC personaController) Borrar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Error al convertir ID", err)
		sError := NewBadRequestError("Error al convertir ID, Borrar")
		c.JSON(http.StatusInternalServerError, sError)
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
		logger.Error("Error al convertir ID, Actualizar", err)
		sError := NewBadRequestError("Error al convertir ID")
		c.JSON(http.StatusInternalServerError, sError)
		return
	}

	var persona domain.Persona

	if err := c.ShouldBindJSON(&persona); err != nil {
		logger.Error("Error en Actualizar persona, parametros", err)
		restErr := NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	persona.ID = id
	p, err := pC.personaService.Actualiza(&persona)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, p)

}

func (pC personaController) ActualizarParcial(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Error en Actualizar Parcial persona, id", err)
		sError := NewBadRequestError("Error ID")
		c.JSON(http.StatusInternalServerError, sError)
		return
	}

	var persona domain.Persona

	if err := c.ShouldBindJSON(&persona); err != nil {
		logger.Error("Error en Actualizar Parcial persona, parametros", err)
		restErr := NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	persona.ID = id
	p, err := pC.personaService.ActualizaParcial(&persona)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, p)

}
