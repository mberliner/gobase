package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mapeaUrls() {
	router.GET("/ping", ping)

	router.POST("/users", userController.Crear)
	router.POST("/users/login", userController.Login)

	router.POST("/personas", personaController.Crear)
	router.GET("/personas/:id", personaController.BuscarPorId)
	router.GET("/personas", personaController.BuscarTodo)
	router.DELETE("/personas/:id", personaController.Borrar)
	router.PUT("/personas/:id", personaController.Actualizar)
	router.PATCH("/personas/:id", personaController.ActualizarParcial)
	/*



		router.GET("/personas/search", personaController.Buscar)

	*/
}

//solo para probar la app
func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
