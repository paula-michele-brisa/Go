package main

import "fmt"

type carro struct {
	nome       string
	velocidade int
}

type ferrari struct {
	carro       // campos anonimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40 carro"
	f.velocidade = 0
	f.turboLigado = true

	fmt.Printf("A ferrai %s está com o turbo ligado? %v\n", f.nome, f.turboLigado)
}
