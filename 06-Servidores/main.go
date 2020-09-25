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
			log.Println(err)
			continue
		}
		go presenta(conn)
	}
}

func presenta(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "Llego: %s\n", ln)
		fmt.Println("\nGoroutine", runtime.NumGoroutine())
	}
	defer conn.Close()

	fmt.Println("Saliendo de esta conexion particular, ya no hay canal")
}
