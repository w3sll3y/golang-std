package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Numeros => ", a, "&", b)
	fmt.Println("Somar => ", a+b)
	fmt.Println("Subtracao => ", a-b)
	fmt.Println("Divisao => ", a/b)
	fmt.Println("Multiplicacao => ", a*b)
	fmt.Println("modulo => ", a%b)

	// bitwise
	fmt.Println("E =>", a&b)   //11 & 10 = 10
	fmt.Println("Ou =>", a|b)  //11 | 10 = 11
	fmt.Println("Xor =>", a^b) //11 ^ 10 = 01

	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	c := 3.0
	d := 2.0
	fmt.Println("Menor =>", math.Min(float64(c), float64(d)))
	fmt.Println("Exponencial =>", math.Pow(c, d))

}
