package main

import (
	"fmt"
)

func main() {

	cPar := make(chan int)
	cImpar := make(chan int)
	cComunica := make(chan bool)

	go enviar(cPar, cImpar, cComunica)

	recibir(cPar, cImpar, cComunica)

	fmt.Println("Fin programa.")
}

func enviar(par chan<- int, impar chan<- int, com chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(com)
}

func recibir(par <-chan int, impar <-chan int, com <-chan bool) {
	for {
		select {
		case p := <-par:
			fmt.Println("Canal Paar:", p)
		case im := <-impar:
			fmt.Println("canal Impar", im)
		case msg, ok := <-com:
			if !ok {
				fmt.Println("Canal Comunicaciones:", msg)
				return
			} else {
				fmt.Println("Canal Comunicaciones", msg)

			}
		}
	}
}
