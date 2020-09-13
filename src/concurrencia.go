package main

import (
	"fmt"
	"runtime"
	"sync"
)

var concurr sync.WaitGroup

type Trabajo struct {
	Contador int
}

type Control struct {
	Bloqueo sync.Mutex
}

var miTrabajo = Trabajo{
	Contador: 0,
}

var miControl = Control{}

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
	fmt.Println("Goroutine", runtime.NumGoroutine(), "Contador: ", miTrabajo.Contador)
}

func reparte() int {
	fmt.Println("Selecciono y envio paquetes a procesar")
	return 100
}

func trabaja(i int) {
	miControl.Bloqueo.Lock()
	v := miTrabajo.Contador
	fmt.Println("Proceso cada paquete por rutinas concurrentes, paquete: ", i)
	runtime.Gosched()
	v++
	miTrabajo.Contador = v
	miControl.Bloqueo.Unlock()
	concurr.Done()
}
