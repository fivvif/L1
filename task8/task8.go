package task8

import "fmt"

func Task8() {
	var num int64 = 11<<22 + 123
	fmt.Println(setBit(num, 0, 0))

}

func setBit(n int64, i uint, bitValue int) int64 {
	if bitValue == 1 {
		// Установить i-й бит в 1
		return n | (1 << i)
	} else {
		// Установить i-й бит в 0
		return n &^ (1 << i)
	}
}
