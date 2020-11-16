package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/controller/rest_errors"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/domain"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/service"
)

type UserController interface {
	Crear(c *gin.Context)
	Login(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(uS service.UserService) UserController {
	return &userController{uS}
}

func (uC userController) Crear(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body " + err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}
	rUser, err := uC.userService.CreaUsuario(&user)
	if err != nil {
		sError := rest_errors.NewInternalServerError("Error al crear usuario", err)
		c.JSON(http.StatusInternalServerError, sError)
		return
	}
	c.JSON(http.StatusCreated, rUser)
}

func (uC userController) Login(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body " + err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}
	pUser, err := uC.userService.Autentica(user.Usuario, user.Password)
	if err != nil {
		restErr := rest_errors.NewInternalServerError("Error al autenticar", err)
		c.JSON(restErr.Status(), restErr)
		return
	}
	c.JSON(http.StatusOK, pUser)
}
