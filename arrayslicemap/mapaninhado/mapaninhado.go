package main

func mian() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 125454.5,
		},
		"J": {
			"José": 12545454,
		},
		"P": {
			"Pedro": 2121212,
		},
	}

	delete(funcsPorLetra, "P")
}
