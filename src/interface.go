package main

import (
	"fmt"
)

//declaraciones siempre con el mismo formato "type identificador tipo"
type agente struct{
	nombre string
	apellido string 
	licencia int
}

type cirujano struct{
	nombre string
	apellido string
	matricula int
}

//Damos comportamiento a los tipos
type human interface{
	speak()
}

func main(){
	juan := cirujano{
		nombre: "Juan",
		apellido: "Rivas",
		matricula: 254321,
	}
	james := agente{
		nombre: "James",
		apellido: "Bond",
		licencia: 007,		
	}
	
	//Usamos comportamiento
	juan.speak()
	james.speak()
	
	//Llamamos desde afuera
	fmt.Println("\nEjemplo Polimorfismo--------------------------")
	quien(juan)
	quien(james)
	
	
}

func (a agente) speak (){
	fmt.Printf("Soy un %T ", a)
	fmt.Println("Me gusta tirar tiros")
}

func (c cirujano) speak (){
	fmt.Printf("Soy un %T\n", c)
	fmt.Println("Me gusta Operar")
}

func quien(h human){
	fmt.Printf("Soy un ---- %T\n", h)
	fmt.Println(h, "Voy a polimorfar")
	h.speak() //Esto es polimorfismo, puedo llamar a los humanos que quieran hablar incluso si hoy otdavia no existen
}