package main

import "fmt"

// nao tem ternario
func obterResultado(nota float64) string {
	// return nota >= 6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		fmt.Println("Aprovado")
	}
	return "Reprovado"
}

func main() {
	obterResultado(6.2)
}
