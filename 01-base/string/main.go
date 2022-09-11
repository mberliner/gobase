package main

import (
	"fmt"
	"strings"
)

func main() {
	frase := "Hola Mundo, esto es un acento: ó"
	fmt.Println(frase)
	fmt.Printf("Tipo de datos: %T\n", frase)

	//Los strings son arrays de bytes
	bs := []byte(frase)

	fmt.Println(bs)
	fmt.Printf("Tipo de dato string interno: %T\n", bs)
	fmt.Printf("Imprimo Unicodes: %#U\n", bs) //Unicode

	for i, l := range bs {
		fmt.Println(i, l)
	}

	fmt.Println("Imprimo UTF-8:")
	for i := 128; i < 200; i++ {
		fmt.Printf("%v - %v - %v - %#U\n", i, rune(i), []byte(string(i)), i)
	}

	//Runes
	fmt.Println("\nMuestro Ejemplo de Runas:")
	fmt.Printf("%v - %v - %v - %#U - %v\n", 'A', string('A'), []byte(string('A')), 'A', rune(65))

	if 'A' == rune('A') {
		fmt.Println("'A' == rune('A')")
	}

	fmt.Println("Buscando en string:", strings.Contains(frase, "ó"))
	if strings.Contains(frase, "ó") {
		fmt.Println("Encontrado ok")
	}

	fraseSeparada := strings.Split(frase, " ")
	fmt.Printf("String inicial: Tipo %T - %v\n", frase, frase)
	fmt.Printf("Separo por blancos: Tipo %T - %v\n", fraseSeparada, fraseSeparada)
	for i, s1 := range fraseSeparada {
		fmt.Println("Indice: ", i, "palabra:", s1)
	}
}
