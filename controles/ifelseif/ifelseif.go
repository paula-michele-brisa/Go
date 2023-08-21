package main

import "fmt"

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 8 && nota < 8 {
		return "B"

	} else if nota >= 3 && nota < 5 {
		return "c"

	} else {
		return "D"

	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))

}
