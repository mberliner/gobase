package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hola Mundo, esto es un acento: ó"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	//Los strings son arrays de bytes
	bs := []byte(s)

	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Printf("%#U\n", bs) //Unicode

	for i, l := range bs {
		fmt.Println(i, l)
	}

	fmt.Println("Imprimo UTF-8:")
	for i := 128; i < 200; i++ {
		fmt.Printf("%v - %v - %v - %#U\n", i, string(i), []byte(string(i)), i)
	}

	//Runes
	fmt.Printf("%v - %v - %v - %#U - %v\n", 'A', string('A'), []byte(string('A')), 'A', rune(65))

	if 'A' == rune('A') {
		fmt.Println("Rune('C') == 'C' ")
	}

	fmt.Println("Buscando en string:", strings.Contains(s, "ó"))
	separado := strings.Split(s, " ")
	fmt.Printf("Separo por blancos: %T - %v\n", separado, separado)
}
