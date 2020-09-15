package main

import (
	"fmt"
)

func main() {

	i := 10
	defer ultima(i) //Asegura que va al final, sirve para cerrar recursos o esperar threads
	//Y evalua sus parámetros al momento de su evaluación (no al final)
	defer cierro(i) //Se ejecutan desde la última hasta la primera
	otra()
	otra()
	otra()

	func() {
		fmt.Println("Funciona anónima sin parametros")
	}()

	func(x int, y int) {
		fmt.Println("Función anónima con parametros: ", x, y)
	}(100, 1000)

	//Asigno una expersion a una variable
	f := func(x int, y int) {
		fmt.Println("Función a una variable con parametros: ", x, y)
	}

	//uso f
	f(200, 2002)
	fmt.Printf("Funcion: - %T\n\n", f)

	//Asigno una funcion que retorna una función
	f2 := bar()
	//Y corre
	fmt.Printf("Funcion2: - %T - resultado: %v\n\n", f2, f2())
	//O lo mismo de otra manera
	fmt.Println("Funcion2:", bar()())

	//Callback
	suma_5 := uso_callback(sum, 3, 4, 5, 6, 7, 8, 9, 10, 15)
	fmt.Println("La suma callback", suma_5)

	//Recursion
	fa := factorial(4)
	fmt.Println("factorial", fa)

	i++ //aunque incremente i la funcion defer no lo tom en cuenta
}

func ultima(i int) {
	fmt.Println("\nEsta va a final de main porque es defer, argumento: ", i)
}

func cierro(i int) {
	fmt.Println("\nEsta va a final de main porque es defer cierro, argumento: ", i)
}

func otra() {
	fmt.Println("esta es otra funcion")
}

func bar() func() int {
	return func() int {
		return 90001
	}
}

func uso_callback(s func(x ...int) int, y ...int) int {
	fmt.Println("Hago inicio -------------")

	var parcial []int
	for _, v := range y {
		if v%5 == 0 {
			parcial = append(parcial, v)
		}
	}
	fmt.Println("Hago fin -------------")
	return s(parcial...)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
