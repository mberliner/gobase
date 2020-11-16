package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/controller/rest_errors"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/domain"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/service"
)

type PersonaController interface {
	BuscarTodo(c *gin.Context)
	Crear(c *gin.Context)
	BuscarPorId(c *gin.Context)
	Borrar(c *gin.Context)
	Actualizar(c *gin.Context)
}

type personaController struct {
	personaService service.PersonaService
}

func NewPersonaController(pS service.PersonaService) PersonaController {
	return &personaController{pS}
}

func (pC personaController) BuscarTodo(c *gin.Context) {

	personas, err := pC.personaService.BuscaTodo()

	if err != nil {
		sError := rest_errors.NewInternalServerError("Error al buscar Personas", err)
		c.JSON(http.StatusInternalServerError, sError)
		return
	}
	c.JSON(http.StatusOK, personas)

}

func (pC personaController) Crear(c *gin.Context) {
	var persona domain.Persona

	if err := c.ShouldBindJSON(&persona); err != nil {
		log.Println("Error en Crear persona:", err, "param:", persona)
		restErr := rest_errors.NewBadRequestError("invalid json body " + err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	p, err := pC.personaService.CreaPersona(&persona)
	if err != nil {
		sError := rest_errors.NewInternalServerError("Error al crear Persona", err)
		c.JSON(http.StatusInternalServerError, sError)
		return
	}

	c.JSON(http.StatusCreated, p)

}

func (pC personaController) BuscarPorId(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		sError := rest_errors.NewBadRequestError("Error al convertir ID" + err.Error())
		c.JSON(http.StatusInternalServerError, sError)
		return
	}

	p, err := pC.personaService.BuscaPersona(id)
	if err != nil {
		sError := rest_errors.NewInternalServerError("Error al buscar Persona por ID", err)
		c.JSON(http.StatusInternalServerError, sError)
		return
	}

	c.JSON(http.StatusCreated, p)

}

func (pC personaController) Borrar(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		sError := rest_errors.NewBadRequestError("Error al convertir ID" + err.Error())
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
	var persona domain.Persona

	if err := c.ShouldBindJSON(&persona); err != nil {
		log.Println("Error en Actualizar persona:", err, "param:", persona)
		restErr := rest_errors.NewBadRequestError("invalid json body " + err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}
	p, err := pC.personaService.Actualiza(&persona)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, p)

}
