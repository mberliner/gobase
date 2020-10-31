package main

import (
	"fmt"
)

func main() {

	i := 10
	defer ultima(i) //Asegura que va al final, sirve para cerrar recursos o esperar threads
	//Y evalua sus parámetros al momento de su evaluación (no al final)
	defer cierro(i) //Se ejecutan desde la última hasta la primera

	func() {
		fmt.Println("Funciona anónima sin parametros")
	}()

	func(x int, y int) {
		fmt.Println("Función anónima con parametros: ", x, y)
	}(100, 1000)

	//Asigno una expresion a una variable
	f := func(x int, y int) {
		fmt.Println("Función asignada a una variable con parametros x int, y int los argumentos aún no están presentes")
	}

	//uso f con argumentos
	f(200, 2002)
	fmt.Printf("Funcion: - %T\n\n", f)

	mantiene := retornaFunc()
	fmt.Println("Funcion que mantiene el valor por scope de variable en main:", mantiene())
	fmt.Println("Funcion que mantiene el valor por scope de variable en main:", mantiene())

	fmt.Println("Funcion que NO mantiene el valor por scope de función retornada:", retornaFunc()())
	fmt.Println("Funcion que NO mantiene el valor por scope de función retornada:", retornaFunc()())

	//Asigno una funcion que retorna una función
	f2 := bar()
	//Y corre
	fmt.Printf("\nFuncion2: - %T - resultado: %v\n\n", f2, f2())
	//Lo mismo de otra manera
	fmt.Println("Funcion2:", bar()())

	//Callback
	suma := usoCallback(sum, 3, 4, 5, 6, 7, 8, 9, 10, 15)
	fmt.Println("La suma callback", suma)

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

func bar() func() int {
	return func() int {
		return 90001
	}
}

func usoCallback(s func(x ...int) int, y ...int) int {
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

func retornaFunc() func() int {

	//Lo importante acá es notar que la variable "cont"
	//se mantiene al ser invocada esta función y alojada en una varaible externa
	cont := 0

	return func() int {
		cont++
		return cont
	}
}
