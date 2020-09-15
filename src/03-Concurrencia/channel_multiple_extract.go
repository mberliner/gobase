package main

import (
	"fmt"
	"sync"
)

func main() {

	cPar := make(chan int)
	cImpar := make(chan int)
	cMux := make(chan int)

	//Envio por canales en rutina separada
	go enviar(cPar, cImpar)

	//Recibo de ambos canales en "paralelo" a uno sólo final (Fan In)
	go recibirEnParalelo(cPar, cImpar, cMux)

	//Extraemos del canal final
	for r := range cMux {
		fmt.Println("Extraido:", r)
	}

	fmt.Println("Fin programa.")
}

func enviar(par chan<- int, impar chan<- int) {
	for i := 0; i < 1000000; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
}

func recibirEnParalelo(par <-chan int, impar <-chan int, mux chan<- int) {
	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		for vo := range impar {
			mux <- vo
		}
		wait.Done()
	}()

	go func() {
		for vp := range par {
			mux <- vp
		}
		wait.Done()
	}()

	wait.Wait()
	close(mux)
}
