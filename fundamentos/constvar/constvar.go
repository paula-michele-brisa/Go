package main

import (
	"fmt"
	"math"
	// letra antes da importacao funciona como as
	// _ indica que o import não será utilizado
)

func main() {

	const PI float64 = 3.1415
	var raio = 3.2

	area := PI * math.Pow(raio, 2)

	fmt.Println("A área da circunferencia é:", area)

	const (
		a = 1
		b = 2
	)

	// var (
	// 	c = 1
	// 	d = 2
	// )

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "etapa"
	fmt.Println(g, h, i)

}
