package main

import "fmt"

func main() {
	fmt.Print("Mesma ")
	fmt.Print("Linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516
	xs := fmt.Sprint(x)
	fmt.Println("o valor de x eh" + xs)
	fmt.Println("o valor de x eh", xs)

	fmt.Printf("o valor de x eh %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d)
	// fmt.Println("\n %v %v %v %v", a, b, c, d) para usar formatado, precisa ser printf = F formated
	// d inteiro
	// f float
	// t boolean
	// 1.f 1 casa decimal
	// s string
	// v valor correto
}
