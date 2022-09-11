package main

import (
	"fmt"
	"time"
)

func main() {
	inicio := time.Now()
	var fechaYHoraPasada time.Time

	fmt.Println("Fecha y Hora actual:", time.Now())
	fmt.Println("Expresado en UTC   :", time.Now().UTC())
	fmt.Println("Fecha y Hora actual con Formato:", time.Now().Format("2006-01-02 15:04:05")) //Formato 01:mes 02:dia 03:hh12 15:HH24 04:min 05:seg 06:año
	fmt.Println()

	fechaYHoraPasada, error := time.Parse("2006-01-02 15:04:05", "2021-09-10 20:32:05")
	if error != nil {
		fmt.Println("Error en conversión:", error)
		fmt.Println()
		return
	}
	fmt.Println("Desde:", fechaYHoraPasada, "Pasaron", (time.Now().Sub(fechaYHoraPasada).Round(time.Hour).Hours())/float64(24), "dias")
	fmt.Println()

	//Ejemplo con error de formato
	fechaYHoraPasada, err := time.Parse("2006-01-02 15:04:05", "1990-12-01 25:01:01")
	if err != nil {
		fmt.Println("Error en conversión:", err)
		fmt.Println()
	}

	fin := time.Now()
	fmt.Println("Duración rutina:", fin.Sub(inicio))
}
