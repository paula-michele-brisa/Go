package main

import "fmt"

func main() {
	i := 1

	// GO não tem artimética de ponteiros
	// * cria o ponteiro
	// ponteiro == referencia de memoria

	var p *int = nil

	p = &i // pegando o endereço da variavel i e apontando para p
	// *p desreferenciando (pegando o valor)

	*p++
	i++

	fmt.Println(p, *p, i, &i)

}
