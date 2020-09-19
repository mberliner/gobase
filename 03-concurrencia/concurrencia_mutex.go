package main

import (
	"fmt"
	"runtime"
	"sync"
)

var concurr sync.WaitGroup

type trabajo struct {
	contador int
}

type control struct {
	bloqueo sync.Mutex
}

var miTrabajo = trabajo{
	contador: 0,
}

var miControl = control{}

func main() {
	fmt.Println("OS       ", runtime.GOOS)
	fmt.Println("ARCH     ", runtime.GOARCH)
	fmt.Println("CPUs     ", runtime.NumCPU())

	fmt.Println("\nGoroutine", runtime.NumGoroutine())

	//Trabajo concurrente con estructuras comunes (race conditions sobre miTrabajo) y mutex para evitarlas
	rp := reparte()
	concurr.Add(rp)

	for i := rp; i > 0; i-- {
		fmt.Println("Paquete: ", i)
		go trabaja(i)
	}

	concurr.Wait()
	fmt.Println("Goroutine", runtime.NumGoroutine())
	fmt.Println("Goroutine", runtime.NumGoroutine(), "Contador: ", miTrabajo.contador)
}

func reparte() int {
	fmt.Println("Selecciono y envio paquetes a procesar")
	return 100
}

func trabaja(i int) {
	miControl.bloqueo.Lock()
	v := miTrabajo.contador
	fmt.Println("Proceso cada paquete por rutinas concurrentes, paquete: ", i)
	runtime.Gosched()
	v++
	miTrabajo.contador = v
	miControl.bloqueo.Unlock()
	concurr.Done()
}
