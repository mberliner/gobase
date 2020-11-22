package app

import (
	"context"
	"io"
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
	puerto  = "8080"
	logHTTP = "../http.log"
)

var (
	router *gin.Engine
	server *http.Server

	userController    controller.UserController
	personaController controller.PersonaController
)

func init() {

	f, err := os.Create(logHTTP)
	if err != nil {
		panic("Error al inciar server, no puedo crear log HTTP" + err.Error())
	}
	// Log a stdout y archivo.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(f, os.Stdout)

	router = gin.Default()

	server = &http.Server{
		Addr:    ":" + puerto,
		Handler: router,
	}

	userController = controller.UserC
	personaController = controller.PersonaC

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

	//canal para señales interrupted(2) y terminated(15)
	cancelar := make(chan os.Signal)
	signal.Notify(cancelar, syscall.SIGINT, syscall.SIGTERM)

	//Bloquea hasta recibir mensaje(señal)
	s := <-cancelar

	logger.Info("Bajando server HTTP recibi señal " + s.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Error con shutdown server:", err)
		panic("Error con shutdown server")
	}

}
