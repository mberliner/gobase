package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/controller"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/logger"
)

const (
	puerto = "8080"
)

var (
	router            *gin.Engine
	userController    controller.UserController
	personaController controller.PersonaController
	server            *http.Server
)

func init() {
	router = gin.Default()
	userController = controller.UserC
	personaController = controller.PersonaC
	server = &http.Server{
		Addr:    ":" + puerto,
		Handler: router,
	}
}

//StartApp Inicia la aplicación, setea el server y las URLs
func StartApp() {
	mapeaUrls()

	//Server inicia en rutina separada
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Error al inicial server HTTP", err)
			panic("Error al inciar server" + err.Error())
		}

	}()

	logger.Info("Iniciado server HTTP")

	//canal para señales sigint y sigterm
	cancelar := make(chan os.Signal)
	signal.Notify(cancelar, syscall.SIGINT, syscall.SIGTERM)

	//Bloquea hasta recibir mensaje(señal)
	s := <-cancelar

	logger.Info("Bajando server HTTP recibi señal " + s.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

}
