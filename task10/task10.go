package task10

import (
	"fmt"
)

func Task10() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	group := make(map[int][]float64)
	for _, value := range temperatures {
		key := int(value / 10)
		key *= 10
		group[key] = append(group[key], value)
	}
	fmt.Println(group)
	for key, value := range group {
		fmt.Println(key, value)
	}

}
