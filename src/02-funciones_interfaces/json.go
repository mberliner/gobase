package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//Si no están en mayuscula no lo codifica a json!!
type persona struct {
	Apellido string
	Nombre   string
	Edad     int
}

//Funciona igual que la estructura de arriba persona si los nombres de campos son iguales al json
//Sino se pueden vincular los nombres asi:
type personaJSON struct {
	Apellido string `json:"ApellidoNuevo"`
	Nombre   string `json:"Nombre"`
	Edad     int    `json:"Edad"`
}

func main() {

	p1 := persona{
		Nombre:   "Fer",
		Apellido: "Czardas",
		Edad:     33,
	}

	p2 := persona{
		Nombre:   "Raul",
		Apellido: "Barolo",
		Edad:     66,
	}

	fmt.Println("Array a Json:\n")

	personas := []persona{p1, p2}
	fmt.Println(personas)

	r, err := json.Marshal(personas)
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println(string(r))

	fmt.Println("Json a Array:\n")
	//De json a texto, notar "ApellidoNuevo"
	jsontext := `[{"ApellidoNuevo":"Czardas","Nombre":"Fer","Edad":33},{"ApellidoNuevo":"Barolo","Nombre":"Raul","Edad":66}]`
	var s []personaJSON

	err = json.Unmarshal([]byte(jsontext), &s)
	if nil != err {
		fmt.Println("Error unmarshal", err)
	}

	fmt.Println("Unmarshal salida", s)

	//Ahora encoding
	fmt.Println("\nEncoding  ------------------")

	//To standard output
	fmt.Println("\nEncoding to Standard Output")
	err = json.NewEncoder(os.Stdout).Encode(jsontext)
	if nil != err {
		fmt.Println("Error Encoder", err)
	}

	fmt.Println("\nEncoding to File")

	filename := "Test.txt"
	file, err := os.Create(filename)
	if nil != err {
		fmt.Println("Error con el archivo:", filename, err)
		os.Exit(10)
	}

	info, _ := file.Stat()
	fmt.Println("Encoding luego Revisar Archivo: ", filename, "Tamaño:", info.Size())
	//A un archivo
	err = json.NewEncoder(file).Encode(jsontext)
	if nil != err {
		fmt.Println("Error Encoder", err)
	}
	err = file.Sync()
	info, err = file.Stat()
	fmt.Println("Encoding luego Revisar Archivo: ", filename, "Tamaño:", info.Size())
	err = file.Close()
	file, err = os.Open(filename)
	if nil != err {
		fmt.Println("Error con el archivo:", filename, err)
		os.Exit(10)
	}

	fmt.Println("\nCopio el archivo a pantalla: ")
	io.Copy(os.Stdout, file)

}
