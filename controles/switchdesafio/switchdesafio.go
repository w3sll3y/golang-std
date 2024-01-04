package main

import "fmt"

func notaParaConceito(n float64) string {
	switch {
	case n > 9 && n < 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	case n >= 0 && n < 3:
		return "E"
	default:
		return "Nota invalida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(8.9))
	fmt.Println(notaParaConceito(7.9))
	fmt.Println(notaParaConceito(3.9))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(-2.1))
}
