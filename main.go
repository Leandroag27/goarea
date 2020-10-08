package goarea

import "math"

// Pi é a relação entre o perímetro de uma circunferência e seu diâmetro
const Pi = 3.1416

// CircArea é responsável em calcular a área de uma circunferência
func CircArea(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// ReactArea é responsável em calcular a área de um retângulo
func ReactArea(base, altura float64) float64 {
	return base * altura
}

// Essa função não é visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
