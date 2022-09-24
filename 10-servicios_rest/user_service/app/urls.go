package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mapeaUrls() {
	router.GET("/ping", ping)

	router.POST("/users", userController.Crear)
	router.POST("/users/login", userController.Login)

}

// solo para probar la app
func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
