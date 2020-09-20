package main

import "fmt"

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

	s = "ddd"
	fmt.Println(s)

	fmt.Println("Imprimo UTF-8:")
	for i := 0; i < 1000; i++ {
		fmt.Printf("%v - %v - %v - %#U\n", i, string(i), []byte(string(i)), i)
	}

	//Runes
	fmt.Printf("%v - %v - %v - %#U - %v\n", 'A', string('A'), []byte(string('A')), 'A', rune(65))

	if 'A' == rune('A') {
		fmt.Println("Rune('C') == 'C' ")
	}

}
