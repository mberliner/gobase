package main

import (
	"fmt"
	"time"
)

func main() {
	var fechaYHoraPasada time.Time
	fmt.Println("Fecha y Hora actual:", time.Now(), "Con UTC", time.Now().UTC())
	fmt.Println("Fecha y Hora actual con Formato:", time.Now().Format("2006-01-02 15:04:05"), "\n")

	fechaYHoraPasada, _ = time.Parse("2006-01-02 15:04:05", "1990-01-01 01:01:01")
	fmt.Println("Desde:", fechaYHoraPasada, "Pasaron", (time.Now().Sub(fechaYHoraPasada).Round(time.Hour).Hours())/float64(24), "dias\n")

	fechaYHoraPasada, err := time.Parse("2006-01-02 15:04:05", "1990-13-01 01:01:01")
	if err != nil {
		fmt.Println("Error en conversión:", err)
	}
}
