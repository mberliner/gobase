package main

import (
	"fmt"
	"encoding/json"
)

//Si no están en mayuscula no lo codifica a json!!
type persona struct{
	Apellido string
	Nombre string
	Edad int
}

//Funciona igual que la estructura de arriba persona si los nombres de campos son iguales al json
//Sino se pueden vincular los nombres asi: 
type personaJson struct{
	Apellido string `json:"ApellidoNuevo"`
	Nombre string `json:"Nombre"`
	Edad int `json:"Edad"`
}


func main(){
	
	p1 := persona{
		Nombre: "Fer",
		Apellido: "Czardas",
		Edad: 33,
	}
	
	p2 := persona{
		Nombre: "Raul",
		Apellido: "Barolo",
		Edad: 66,
	}
	
	
	e := []persona{p1,p2,}
	fmt.Println(e)
	
	r,err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(r))
	
	//De json a texto, notar ApellidoNuevo
	json1 := `[{"ApellidoNuevo":"Czardas","Nombre":"Fer","Edad":33},{"ApellidoNuevo":"Barolo","Nombre":"Raul","Edad":66}]`
	var s  [] personaJson
	
	err = json.Unmarshal([]byte(json1), &s)
	if err != nil {
		fmt.Println("Error unmarshal", err)
	}
	
	fmt.Println("Unmarshal salida", s)
}