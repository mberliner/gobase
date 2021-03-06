package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://rae.org",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://cont.ar",
	}

	canal := make(chan string)

	//La primera vez busco la respuesta de las URLs y agrego al canal
	for _, sitio := range sites {
		go verificaSitio(sitio, canal)
	}

	//Luego recorro el canal de forma infinita
	for sitiotemp := range canal {

		go func(sitio string) {
			//function literal, va con parametros porque sino estaria usando el string de main
			//y cambia sin control desde afuera y lo usa la rutina

			//Sleep para no llamar tan seguido a las verificaciones, si estuviese fuera de una func literal
			// deberia esperar a la respuesta de cada sitio antes de llamar al siguiente
			time.Sleep(5 * time.Second)
			verificaSitio(sitio, canal)
		}(sitiotemp)

	}

}

func verificaSitio(sitio string, canal chan string) {
	resp, error := http.Get(sitio)

	if error != nil {
		fmt.Println(sitio, "no funciona:")
		canal <- sitio
		return
	}

	fmt.Println(sitio, "is ok!")
	fmt.Println(resp, "\n")

	//Agrega y bloquea
	canal <- sitio

}
