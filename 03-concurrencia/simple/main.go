package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var concurr sync.WaitGroup

func main() {
	fmt.Println("OS       ", runtime.GOOS)
	fmt.Println("ARCH     ", runtime.GOARCH)
	fmt.Println("CPUs     ", runtime.NumCPU())

	fmt.Println("\nGoroutine", runtime.NumGoroutine())

	//Trabajo concurrente a la vieja usanza sin estructuras comunes (sin race conditions)
	rp := reparte()
	concurr.Add(rp)

	for i := rp; i > 0; i-- {
		fmt.Println("Paquete: ", i)
		go trabaja(i)
	}
	fmt.Println("Goroutine", runtime.NumGoroutine())

	concurr.Wait()
	fmt.Println("Goroutine", runtime.NumGoroutine())

}

func reparte() int {
	fmt.Println("Selecciono y envio paquetes a procesar")
	return 100
}

func trabaja(i int) {
	fmt.Println("Proceso cada paquete por rutinas concurrentes, paquete: ", i, "ID Goroutine: ", getGID())
	runtime.Gosched()
	concurr.Done()
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
