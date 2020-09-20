package main

import (
	"fmt"
)

func main() {

	canal := make(chan int)
	fmt.Printf("%T\n", canal)

	go func() {
		for v := 0; v < 100; v++ {
			canal <- v
		}
		close(canal)
	}()

	//Lee hasta que el canal esté cerrado
	for r := range canal {
		fmt.Println("Recibido:", r)
	}

	fmt.Println("Fin programa.")
}
