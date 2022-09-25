package main

import (
	"fmt"
	"runtime"
	"sync"
)

var trabajadores sync.WaitGroup

type trabajo struct {
	contador int
	etiqueta []string
}

type sincro struct {
	bloqueo sync.Mutex
}

var miTrabajo = trabajo{
	contador: 0,
	etiqueta: []string{},
}

var sincroTrabajo = sincro{}

func main() {
	fmt.Println("OS       ", runtime.GOOS)
	fmt.Println("ARCH     ", runtime.GOARCH)
	fmt.Println("CPUs     ", runtime.NumCPU())

	fmt.Println("\nGoroutine", runtime.NumGoroutine())

	//Trabajo concurrente con estructuras comunes (race conditions sobre miTrabajo) y mutex para evitarlas
	cantTrabajo := reparte(10)
	trabajadores.Add(cantTrabajo)

	fmt.Println("Comienza el trabajo: ", miTrabajo)
	for i := 0; i < cantTrabajo; i++ {
		fmt.Println("Paquete: ", i)
		go trabaja(i)
	}

	trabajadores.Wait()
	fmt.Println("Goroutine", runtime.NumGoroutine())
	fmt.Println("Goroutine", runtime.NumGoroutine(), "Contador: ", miTrabajo.contador)
}

func reparte(cant int) int {
	fmt.Println("Selecciono y envio ", cant, " paquetes a procesar")
	var i int
	for i = 0; i < cant; i++ {
		miTrabajo.etiqueta = append(miTrabajo.etiqueta, fmt.Sprint("string ----> ", i))
	}
	return i
}

func trabaja(i int) {
	sincroTrabajo.bloqueo.Lock()
	v := miTrabajo.contador
	fmt.Println("Proceso cada paquete por rutinas concurrentes, paquete: ", i, "mitrabajo: ", miTrabajo.etiqueta[i])
	runtime.Gosched()
	v++
	miTrabajo.contador = v
	sincroTrabajo.bloqueo.Unlock()
	trabajadores.Done()
}
