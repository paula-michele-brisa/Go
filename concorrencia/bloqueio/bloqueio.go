package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operacao bloqueante
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) // canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // operacao bloqueante
	fmt.Println("Foi  lido")
	fmt.Println(<-c)    // deadlock
	fmt.Println("Fim!") // nunca será executado porque na linha anterior foi gerado um deadlock
}
