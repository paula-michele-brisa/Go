package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // o compilador vai contar

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// obtendo o numero
	for _, numero := range numeros {
		fmt.Println(numero)
	}

	// obtendo apenas o indice
	for indice := range numeros {
		fmt.Println(indice)
	}

}
