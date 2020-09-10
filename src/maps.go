package main

import (
	"fmt"
)

func main() {

	//Los mapas estàn indexados internamente y no tienen orden
	cosas := map[string] string{
		"primera" : "cosa1",
		"segunda" : "cosa2",
	}
	
	var otromapa  map [string] string
	
	
	mapin := make(map[int] string)  //Buil in function make
	
	fmt.Println(cosas)
	
	//Si pido una key que no existe da cero (si es int) o string vacio y no da error
	fmt.Println("Si no existe da vacio: ", cosas["no existe"])
	
	k, v := cosas["no existe"]
	fmt.Println("Si no existe da vacio y boleano: ", k, v)
	
	//Elegante y simple forma de valida si existe
	if k, ok := cosas["no existe"] ; ok {
		fmt.Println("Existe: ", k, ok)
	} else {
		fmt.Println("No Existe: ", k)
	}
	
	//Recorro
	for key,value := range cosas{
		fmt.Println("Clave:", key, " - Valor:", value)
	}
		
	fmt.Println(otromapa)
	
	mapin[123] = "Salga de aqui"
	fmt.Println(mapin)
	
	delete(mapin, 123) //Built in para borrar
		//Si borro lo que no existe no da error
	delete(mapin, 1)
	fmt.Println(mapin)
	
		
		
}