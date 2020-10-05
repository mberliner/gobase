package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	//El certificado cert.pem se obtiene de una autoridad CA
	//letsencrypt.org en una CA y los provee de forma gratuita
	//Tiene Apis para negociarlos de forma automática
	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func foo(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Ejemplo sin certificado extendido por autoridad.\n"))
}
