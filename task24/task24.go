package task24

import (
	"fmt"
	"math"
)

// Point представляет точку с координатами x и y.
type Point struct {
	x float64
	y float64
}

// NewPoint конструктор
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// FindRange вычисляет расстояние между двумя точками p1 и p2 с использованием
// теоремы Пифагора и возвращает результат в виде float64.
func FindRange(p1, p2 *Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func Task24() {
	// создаём 2 точки
	p1 := NewPoint(6, 2)
	p2 := NewPoint(4, 2)
	fmt.Println(FindRange(p1, p2))
}
