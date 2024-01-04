package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro eh", reflect.TypeOf(32000))

	// sem sinal (so positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Print("byte eh ", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("o valor maximo do int eh ", i1)
	fmt.Println("o valor de i1 eh ", reflect.TypeOf(i1))

	var i2 rune = 'a'
	fmt.Println("o tipo do rune eh ", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("o tipo de x eh ", reflect.TypeOf(x))
	fmt.Println("o tipo de literal de 49.99 eh ", reflect.TypeOf(49.99)) //float 64

	// boolean
	bo := true
	fmt.Println("o tipo de bo eh ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Ola meu nome eh Wesley"
	fmt.Println(s1 + "!")
	fmt.Println("o tamanho da string eh ", len(s1))
	fmt.Println("o tipo de s1 eh ", reflect.TypeOf(s1))

	// string com multiplas linhas
	s2 := `Ola
	meu 
	nome 
	eh
	Wesley`
	fmt.Println("o tamanho da string s2 eh ", len(s2))

	// char??
	// var x char = 'a' nao existe
	char := 'a'
	fmt.Println("o tipo de char eh ", reflect.TypeOf(char)) //int32
	fmt.Println(char)
}
