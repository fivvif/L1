package task22

import "fmt"

func Task22() {
	a := 6 << 20
	b := 5 << 20

	fmt.Println(a + b)
	fmt.Println(a * b)
	fmt.Println(a - b)
	aFloat := float64(a)
	bFloat := float64(b)
	fmt.Println(aFloat / bFloat)
}
