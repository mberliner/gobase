package main

import "fmt"

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int)
	go genera(ch)

	for i := 1; i <= 10; i++ {
		primo := <-ch
		fmt.Println("Numero Primo", i, "es:", primo)
		ch1 := make(chan int)
		go filtra(ch, ch1, primo)
		ch = ch1
	}
}

func genera(ch chan<- int) {
	for i := 2; ; i++ {
		fmt.Println("Genera: ", i)
		ch <- i
	}
}

// Copia del channel 'in' al 'out', elimina los divisibles por primo
func filtra(in <-chan int, out chan<- int, primo int) {
	for {
		i := <-in
		fmt.Println("Filtra: ", i, primo, i%primo)
		if i%primo != 0 {
			out <- i
		}
	}
}
