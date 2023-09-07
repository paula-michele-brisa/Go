package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre as goroutines
// Channel é um tipo

func multiplica(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base

}

func main() {
	c := make(chan int)
	go multiplica(2, c)
	fmt.Println("Continua executando")

	a, b := <-c, <-c // recebendo os dados do canal
	fmt.Println("Fez uma pausa para receber os valores a partir do canal")
	fmt.Println(a, b)

	fmt.Println(<-c)
}
