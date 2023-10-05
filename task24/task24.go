package task24

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func FindRange(p1, p2 *Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func Task24() {
	p1 := NewPoint(6, 2)
	p2 := NewPoint(4, 2)
	fmt.Println(FindRange(p1, p2))
}
