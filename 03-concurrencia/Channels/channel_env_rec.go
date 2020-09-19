package main

import (
	"fmt"
)

func main() {

	//Creo un canal y luego lo uso como envio y recepcion por separado
	canal := make(chan int)
	fmt.Printf("%T\n", canal)

	for i := 0; i < 100; i++ {
		go envio(canal, i) //Desperdicio de mem de stack
		fmt.Println("enviado:", i)
	}

	//Si va mas allá de 100 bloquea el programa!!!
	for i := 0; i < 100; i++ {
		r := recibo(canal)
		fmt.Println("Recibido:", r)
	}
}

func envio(c chan<- int, v int) {
	c <- v
}

func recibo(c <-chan int) int {
	return (<-c)
}
