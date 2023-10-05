package task13

import "fmt"

func Task13() {
	a := 333
	b := 111
	fmt.Println(a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}

func Task13_2() {
	a := 333
	b := 111
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
