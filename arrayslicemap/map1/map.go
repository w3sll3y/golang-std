package main

import "fmt"

func main() {
	// var aprovados map[int]string
	//maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[21234567890] = "Maria"
	aprovados[31231312312] = "Pedro"
	aprovados[29809809877] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[21234567890])
	delete(aprovados, 29809809877)
	fmt.Println(aprovados[29809809877])
	fmt.Println(aprovados)

}
