package main

import (
	"fmt"
	"runtime"
	"sync"
)

var concurr sync.WaitGroup

func main(){
	fmt.Println("OS       ", runtime.GOOS)
	fmt.Println("ARCH     ", runtime.GOARCH)
	fmt.Println("CPUs     ", runtime.NumCPU())
	
	
	fmt.Println("\nGoroutine", runtime.NumGoroutine())
	
	//Trabajo concurrente a la vieja usanza sin estructuras comunes (sin race conditions)
	rp := reparte()
	concurr.Add(rp)
	
	for i:=rp; i>0 ;i-- {
		fmt.Println("Paquete: ", i)
		go trabaja(i)
	}
	fmt.Println("Goroutine", runtime.NumGoroutine())
	
	concurr.Wait()
	fmt.Println("Goroutine", runtime.NumGoroutine())
	
}

func reparte() int{
	fmt.Println("Selecciono y envio paquetes a procesar")
	return 100
}

func trabaja(i int){
	fmt.Println("Proceso cada paquete por rutinas concurrentes, paquete: ", i)
	runtime.Gosched()
	concurr.Done()
}