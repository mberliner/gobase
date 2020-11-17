package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/controller"
)

var (
	router            *gin.Engine
	userController    controller.UserController
	personaController controller.PersonaController
)

func init() {
	router = gin.Default()
	userController = controller.UserC
	personaController = controller.PersonaC
}

//StartApp Inicia la aplicación, setea el server y las URLs
func StartApp() {
	mapUrls()
	if err := router.Run(":8080"); err != nil {
		panic("Error al inciar server" + err.Error())
	}
}

func mapUrls() {
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
