package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var trabajadores sync.WaitGroup
var trabajo []string

func main() {
	fmt.Println("OS       ", runtime.GOOS)
	fmt.Println("ARCH     ", runtime.GOARCH)
	fmt.Println("CPUs     ", runtime.NumCPU())

	fmt.Println("\nGoroutine", runtime.NumGoroutine())

	//Trabajo concurrente a la vieja usanza sin estructuras comunes (sin race conditions)
	cantTrabajo := reparte(10)
	trabajadores.Add(cantTrabajo)

	for i := 0; i < cantTrabajo; i++ {
		fmt.Println("Paquete: ", i)
		go trabaja(i)
	}
	fmt.Println("Goroutines al fin envio trabajo: ", runtime.NumGoroutine())

	trabajadores.Wait()
	fmt.Println("Goroutine Final", runtime.NumGoroutine())

}

func reparte(cant int) int {
	fmt.Println("Selecciono y envio ", cant, " paquetes a procesar")
	var i int
	for i = 0; i < cant; i++ {
		trabajo = append(trabajo, fmt.Sprint("string ----> ", i))
	}
	return i
}

func trabaja(paq int) {
	fmt.Println("Proceso cada paquete por rutinas concurrentes, paquete: ", trabajo[paq], "ID Goroutine: ", getGID())
	runtime.Gosched()
	//Decrementa cant trbajos hechos
	trabajadores.Done()
}

// Solo para saber que id tiene la goroutine
// es completamente innecesario
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
