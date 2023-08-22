package main

func main() {
	funcESalarios := map[string]float64{
		"José":  1252.25,
		"Maria": 20000.00,
	}

	delete(funcESalarios, "Inexistente")

}
