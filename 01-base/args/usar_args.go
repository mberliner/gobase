package main

import (
	"flag"
	"fmt"
	_ "log"
)

//Ejemplo para usar ok
//go run usar_args.go -arg1=1 -arg2=estea -arg3=false -arg4=100

//Con error:
//go run usar_args.go -arg1=x -arg2=estea -arg3=false -arg4=100
func main() {
	//Son punteros
	arg1 := flag.Int("arg1", 100, "Primer parámetro")
	arg2 := flag.String("arg2", "Default string", "2do parámetro")
	arg3 := flag.Bool("arg3", true, "3er parámetro")

	//Otra forma usando la variable en lugar del puntero
	var arg4 int
	flag.IntVar(&arg4, "arg4", 10001, "Cuarto parámetro")

	//Sólo para verificar cantidad min de obligatorios, sino usa default
	//	if flag.NArg() < 4 {
	//		log.Fatalln("Debe haber 4 parámetros obligatorios")
	//	}

	flag.Parse()

	fmt.Println("arg1:", *arg1)
	fmt.Println("arg2:", *arg2)
	fmt.Println("arg3:", *arg3)
	fmt.Println("arg4:", arg4)
}
