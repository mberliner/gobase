package app

import (
	"log"
	"net/http"
)

// StartApp Inicia toda la aplicación
func StartApp() {
	mapUrls()
	iniciaServerHTTP()
}

func iniciaServerHTTP() {

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error: ", err)
	}
}
