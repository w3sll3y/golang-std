package main

import "fmt"

func obterValorAprovadoComDefer(numero int) int {
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	} else {
		fmt.Println("Valor baixo")
		return numero
	}
}

func obterValorAprovadoSemDefer(numero int) int {
	fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	} else {
		fmt.Println("Valor baixo")
		return numero
	}
}

func main() {
	fmt.Println("com defer =>")
	fmt.Println(obterValorAprovadoComDefer(6000))
	fmt.Println(obterValorAprovadoComDefer(2000))
	fmt.Println("sem defer =>")
	fmt.Println(obterValorAprovadoSemDefer(6000))
	fmt.Println(obterValorAprovadoSemDefer(2000))
}
