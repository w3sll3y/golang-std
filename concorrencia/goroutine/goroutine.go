package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interacoes %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "PQ vc n fala cmg?", 3)
	// fale("Joao", "So posso falar dps d vc", 1)
	// go fale("Maria", "Eii...", 500)
	// go fale("Joao", "Opaa...", 500)
	// time.Sleep(time.Second * 5)
	// fmt.Println("Fim!")

	go fale("Maria", "Entendii...", 10)
	fale("Joao", "Parabens...", 5)

}
