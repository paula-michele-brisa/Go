package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Metodo: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto

	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	// Outra forma de preencher valores de uma struct
	produto2 := produto{"Notebook", 2.00, 0.5}

	fmt.Println(produto1, produto1.precoComDesconto())
	fmt.Println(produto2, produto2.precoComDesconto())

}
