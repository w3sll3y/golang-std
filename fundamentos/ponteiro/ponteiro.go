package main

import "fmt"

func main() {
	i := 1 // 8 bytes da memoria

	var p *int = nil

	p = &i // pegando endereco da variavel
	*p++   // desreferenciando (pegando valor)
	i++
	// Go nao tem aritimetica de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
