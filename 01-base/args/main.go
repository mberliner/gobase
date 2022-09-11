package main

import (
	"flag"
	"fmt"
	_ "log"
)

//Ejemplo para usar ok
//go run usar_args.go -arg1=1 -arg2=estea -arg3=false -arg4=100

// Con error:
// go run usar_args.go -arg1=x -arg2=estea -arg3=false -arg4=100
func main() {
	//Son punteros
	arg1 := flag.Int("arg1", 100, "Primer parámetro")
	arg2 := flag.String("arg2", "Default string", "2do parámetro")
	arg3 := flag.Float64("arg3", 99.0, "3er parámetro")

	//Otra forma usando la variable en lugar del puntero
	var arg4 int
	flag.IntVar(&arg4, "arg4", 10001, "Cuarto parámetro")

	//Sólo para verificar cantidad min de obligatorios, sino usa default
	//	if flag.NArg() < 4 {
	//		log.Fatalln("Debe haber 4 parámetros obligatorios")
	//	}

	flag.Parse()

	fmt.Printf("arg1: %v, tipo: %T\n", *arg1, *arg1)
	fmt.Printf("arg1: %v, tipo: %T\n", *arg2, *arg2)
	fmt.Printf("arg1: %v, tipo: %T\n", *arg3, *arg3)
	fmt.Printf("arg1: %v, tipo: %T\n", arg4, arg4)
}
