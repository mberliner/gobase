package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	_ "time"
)

var concurr sync.WaitGroup

type Trabajo struct{
	Contador int64
}

type Control struct{
	Bloqueo int64
}

var miTrabajo = Trabajo{
	Contador: 0,
}


func main(){
	fmt.Println("OS       ", runtime.GOOS)
	fmt.Println("ARCH     ", runtime.GOARCH)
	fmt.Println("CPUs     ", runtime.NumCPU())
	
	
	fmt.Println("\nGoroutine", runtime.NumGoroutine())
	
	//Trabajo concurrente con estructuras comunes (race conditions sobre miTrabajo) y contador atomic para evitarlas
	rp := reparte()
	concurr.Add(rp)
	
	for i:=rp; i>0 ;i-- {
		fmt.Println("Paquete: ", i)
		go trabaja(i)
		fmt.Println("Goroutine", runtime.NumGoroutine(), "Contador Concurrente: ", atomic.LoadInt64(&miTrabajo.Contador)) //Concurrencia aqui de lectura
	}
		
	concurr.Wait()
	fmt.Println("Goroutine", runtime.NumGoroutine())
	fmt.Println("Goroutine", runtime.NumGoroutine(), "Contador: ", miTrabajo.Contador) //ya no hay concurrencia aqui
}

func reparte() int{
	fmt.Println("Selecciono y envio paquetes a procesar")
	return 100
}

func trabaja(i int){
	fmt.Println("Proceso cada paquete por rutinas concurrentes, paquete: ", i)
	//time.Sleep(2*time.Second)
	runtime.Gosched()
	atomic.AddInt64(&miTrabajo.Contador, 1)
	
	concurr.Done()
}