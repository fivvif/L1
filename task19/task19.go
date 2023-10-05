package task19

import "fmt"

func Task19() {
	str := "балабол"
	fmt.Println(reverseString(str))
}

func reverseString(original string) string {
	runes := []rune(original)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
