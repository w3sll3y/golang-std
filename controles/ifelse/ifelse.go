package main

import "fmt"

func imprimirNota(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota:", nota)
	} else {
		fmt.Println("Reprovado com nota:", nota)
	}
}

func main() {
	imprimirNota(7.3)
	imprimirNota(5.1)
}
