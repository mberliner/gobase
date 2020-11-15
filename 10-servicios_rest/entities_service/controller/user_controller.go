package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/controller/rest_errors"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/domain"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/service"
)

func getUserId(userIdParam string) (int64, rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Crear(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body " + err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}
	rUser, err := service.UserS.CreaUsuario(&user)
	if err != nil {
		sError := rest_errors.NewInternalServerError("Error al crear usuario", err)
		c.JSON(http.StatusInternalServerError, sError)
		return
	}
	c.JSON(http.StatusCreated, rUser)
}

func Login(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body " + err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}
	pUser, err := service.UserS.Autentica(user.Usuario, user.Password)
	if err != nil {
		restErr := rest_errors.NewInternalServerError("Error al autenticar", err)
		c.JSON(restErr.Status(), restErr)
		return
	}
	c.JSON(http.StatusOK, pUser)
}
