package main

import "fmt"

type esportivo interface {
	ligarTurboo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurboo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurboo()

	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurboo()

	fmt.Println(carro1, carro2)
}
