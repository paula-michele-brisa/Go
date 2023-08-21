package main

import "fmt"

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}

	return "Reprovado"
}

// Não tem operador ternário em GO
func main() {
	fmt.Println(obterResultado(6.2))

}
