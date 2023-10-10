package task21

import (
	"fmt"
	"math"
)

// LegacyPrinter представляет старый принтер с устаревшим интерфейсом.
type LegacyPrinter struct{}

// PrintLegacy выполняет печать с использованием старого принтера.
func (lp *LegacyPrinter) PrintLegacy(message string) {
	fmt.Println("Legacy Printer: " + message)
}

// Printer представляет современный интерфейс принтера.
type Printer interface {
	Print(message string)
}

// ModernPrinter представляет современный принтер.
type ModernPrinter struct{}

// Print выполняет печать с использованием современного принтера.
func (mp *ModernPrinter) Print(message string) {
	fmt.Println("Modern Printer: " + message)
}

// LegacyPrinterAdapter - это адаптер между старым и современным принтерами.
type LegacyPrinterAdapter struct {
	legacyPrinter *LegacyPrinter
}

// Print выполняет адаптированную печать, используя старый принтер через адаптер.
func (adapter *LegacyPrinterAdapter) Print(message string) {
	adapter.legacyPrinter.PrintLegacy(message)
}

func Task21() {
	modernPrinter := &ModernPrinter{}
	legacyPrinter := &LegacyPrinter{}

	// Используем современный принтер
	modernPrinter.Print("Hello, modern world!")

	// Используем адаптер, чтобы использовать старый принтер как современный
	adapter := &LegacyPrinterAdapter{legacyPrinter}
	adapter.Print("Hello, legacy world!")
}

//task21_2

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

//предположим что мы не может менять triangle

type Triangle struct {
	Base   float64
	Height float64
}

type TriangleAdapter struct {
	Triangle
}

func (ta *TriangleAdapter) Area() float64 {
	return 0.5 * ta.Base * ta.Height
}

func (ta *TriangleAdapter) Perimeter() float64 {
	// Для простоты предположим, что у нас есть доступ к функции для расчета длин сторон треугольника.
	side1 := 3.0
	side2 := 4.0
	hypotenuse := math.Sqrt(side1*side1 + side2*side2)
	return side1 + side2 + hypotenuse
}

func Task21_2() {
	circle := &Circle{Radius: 5}
	rectangle := &Rectangle{Width: 3, Height: 4}
	triangle := &Triangle{Base: 3, Height: 4}

	shapes := []Shape{circle, rectangle, &TriangleAdapter{*triangle}}

	for _, shape := range shapes {
		fmt.Printf("Shape Type: %T\n", shape)
		fmt.Printf("Area: %f\n", shape.Area())
		fmt.Printf("Perimeter: %f\n", shape.Perimeter())
		fmt.Println()
	}
}
