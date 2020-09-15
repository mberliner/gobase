package main

import (
	"fmt"
)

func main() {

	//Creo un canal con buffer 2
	//Si supera el buffer bloquea la ejecución
	//Por eso los canales siempre deben ser consumidos desde go rutines
	canal := make(chan int, 2)
	fmt.Printf("%T\n", canal)

	canal <- 100
	canal <- 101

	fmt.Println("Recibido:", <-canal)
	fmt.Println("Recibido:", <-canal)
	//fmt.Println("Recibido:", <-canal) //el 3ro bloquea el programa entero

	//Canal para envios solamente
	canalEnvio := make(chan<- int, 2) //Recepcion make(<- canal)
	fmt.Printf("\n%T solo para envios\n\n", canalEnvio)
	canalEnvio <- 100000

	//s,e:=fmt.Println("Recibido:", <-canalEnvio) //Error es de recepcion solamente

	//Canal sin buffer para recibir desde go rutine
	canalN := make(chan int)

	go func(msg int) {
		canalN <- msg
	}(10001)

	fmt.Println("Recibido desde Go Rutine:", <-canalN)

	//Los canales son tipos asi que se puede hacer conversión (siempre desde lo gral a lo particular)

	canalNEnvio := (chan<- int)(canalN)
	go func(msg int) {
		canalNEnvio <- msg
	}(1111111)

	fmt.Printf("Canal convertido de: %T a: %T\n", canalN, canalNEnvio)

}
