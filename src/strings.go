package main

import "fmt"

func main() {
	s := "Hola Mundo, esto es un acento: ó"
	fmt.Println(s)
	fmt.Printf("%T\n",s)
	
	//Los strings son arrays de bytes
	bs := []byte(s)
	
	fmt.Println(bs)
	fmt.Printf("%T\n",bs)
	fmt.Printf("%#U\n",bs) //Unicode
	
	for i, l := range bs {
		fmt.Println(i,l)
	}
	
	s = "ddd"
	fmt.Println(s)
	
}