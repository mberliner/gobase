package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("CPUs     ", runtime.NumCPU())

	trabajo := make(chan int)
	resultado := make(chan int)

	//Cargo una cantidad de trabajo en un canal
	go reparte(trabajo)

	//Trabajo concurrente resuelto abriendo canales en "paralelo" para resolverlo  (fan out)
	//y luego lo integra en un canal de salida (fan in)
	go trabajarParalelo(trabajo, resultado)

	for r := range resultado {
		fmt.Println("Resultado: ", r)
	}
	fmt.Println("Goroutines del final", runtime.NumGoroutine())

}

func reparte(t chan<- int) {
	for i := 0; i < 100; i++ {
		t <- i
	}
	close(t)
}

func trabajarParalelo(trabajo <-chan int, resultado chan<- int) {
	var wg sync.WaitGroup
	const maxConcurrencia = 10
	wg.Add(maxConcurrencia)
	fmt.Println("Voy a separar el trabajo en", maxConcurrencia, "partes")
	for i := 0; i < maxConcurrencia; i++ {
		go func() {
			for t := range trabajo {
				resultado <- trabajar(t)
			}
			wg.Done()
		}()
	}
	fmt.Println("Goroutines", runtime.NumGoroutine())
	wg.Wait()
	close(resultado)
}

func trabajar(id int) int {
	time.Sleep(2 * time.Second) //cada trabajo demora 2 segundos
	return id * 99
}
