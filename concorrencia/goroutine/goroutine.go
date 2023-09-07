package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Poruqe voce não fala comigo?", 3)
	// fale("João", "Só posso falar depois de voce!", 1)

	// go fale("Maria", "Ei.....", 1000)
	// go fale("João", "Opa.....", 1000)

	go fale("Maria", "Entendi", 10)
	fale("João", "Parabéns!", 5)
}
