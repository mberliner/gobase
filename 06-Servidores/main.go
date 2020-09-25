package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"runtime"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println("Hubo error, pero continuo", err)
			continue
		}
		fmt.Println("Local Adrress:", conn.LocalAddr(), "Remote Address:", conn.RemoteAddr())
		go comunica(conn)
	}

	fmt.Println("Saliendo del programa")
}

//Una por cada conexion aceptada
func comunica(conn net.Conn) {
	fmt.Println("\nNro of Goroutines", runtime.NumGoroutine())
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "Soy tu server y me llego: %s\n", ln)
	}

	fmt.Println("Saliendo de esta conexion particular, ya no hay canal")
}
