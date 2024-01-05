package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15231.70,
			"Guga Pereira":   12222.56,
		},
		"J": {
			"Jose Joao": 13132.34,
		},
		"P": {
			"Pedro Jr": 12312.22,
		},
	}

	fmt.Println(funcsPorLetra)
	delete(funcsPorLetra, "P")
	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

}
