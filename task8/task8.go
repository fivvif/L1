package task8

import "fmt"

func Task8() {
	var num int64 = 11<<22 + 123
	fmt.Println(setBit(num, 0, 0))

}

func setBit(num int64, i uint, bitValue int) int64 {
	fmt.Printf("%b\n", num)
	if bitValue == 1 {
		num |= 1 << i
	} else if bitValue == 0 {
		num &= ^(1 << i)
	}
	fmt.Printf("%b\n", num)
	return num
}
