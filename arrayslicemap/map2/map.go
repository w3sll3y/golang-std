package main

import "fmt"

func main() {
	funcsEsalario := map[string]float64{
		"Jose Joao":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Jr":       1200.0,
	}

	funcsEsalario["Rafael Filho"] = 1350.0
	delete(funcsEsalario, "qweqew")

	for nome, salario := range funcsEsalario {
		fmt.Println(nome, salario)
	}
}
