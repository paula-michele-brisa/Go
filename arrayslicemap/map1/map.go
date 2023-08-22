package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicilizados

	aprovados := make(map[int]string)

	aprovados[1212121] = "Maria"
	aprovados[123] = "Pedro"
	aprovados[234] = "Ana"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[123])

	aprovados[123] = "João"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

}
