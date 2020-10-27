package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func StartApp() {
	mapUrls()
	if err := router.Run(":8080"); err != nil {
		panic("Error al inciar server" + err.Error())
	}
}

func mapUrls() {
	router.GET("/ping", ping)
	/*
		router.POST("/users", users.Create)
		router.GET("/users/:user_id", users.Get)
		router.PUT("/users/:user_id", users.Actualizar)
		router.PATCH("/users/:user_id", users.ActualizarParcial)
		router.DELETE("/users/:user_id", users.Borrar)
		router.GET("/internal/users/search", users.Buscar)
		router.POST("/users/login", users.Login)
	*/
}

//solo para probar la app
func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
