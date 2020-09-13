package main

import (
	"fmt"
	"runtime"
	"sync"
	_ "time"
)

var concurr sync.WaitGroup

type Actualizar struct{
	Contador int
}
var aTrabajar = Actualizar{
		Contador: 0,
}

func main(){
	fmt.Println("OS       ", runtime.GOOS)
	fmt.Println("ARCH     ", runtime.GOARCH)
	fmt.Println("CPUs     ", runtime.NumCPU())
	
	
	fmt.Println("\nGoroutine", runtime.NumGoroutine())
	
	//Trabajo concurrente con estructuras comunes (race conditions sobre aTrabajar)
	rp := reparte()
	concurr.Add(rp)
	
	for i:=rp; i>0 ;i-- {
		fmt.Println("Paquete: ", i)
		go trabaja(i)
	}
	fmt.Println("Goroutine", runtime.NumGoroutine(), "Contador: ", aTrabajar.Contador)
	
	concurr.Wait()
	fmt.Println("Goroutine", runtime.NumGoroutine())
		fmt.Println("Goroutine", runtime.NumGoroutine(), "Contador: ", aTrabajar.Contador)
}

func reparte() int{
	fmt.Println("Selecciono y envio paquetes a procesar")
	return 20
}

func trabaja(i int){
	v := aTrabajar.Contador
	fmt.Println("Proceso cada paquete por rutinas concurrentes, paquete: ", i)
	//time.Sleep(2*time.Second)
	runtime.Gosched()
	v++
	aTrabajar.Contador = v
	concurr.Done()
}