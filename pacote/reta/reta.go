package main

import "math"

// Inicializando com letra maiuscula 'e Publica (visivel dentro e fora do pacote)!
// Inicializando com letra minuscula 'e Privado (visivel apenas dentro do pacote)!

// Por exemplo
// Ponto - gerara' algo publico
// ponto ou _Ponto - gerara' algo privado

// Ponto representa uma cordenada do plano carteziano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia 'e responsavel por calcular a distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
