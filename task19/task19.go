package task19

import "fmt"

func Task19() {
	str := "балабол"
	fmt.Println(reverseString(str))
}

func reverseString(original string) string {
	// Преобразуем строку в массив рун
	runes := []rune(original)
	// последовательно меняем местами элементы по индексам i, j
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
	// Т.к символы могут быть в unicode используем rune, для хранения символов.
}
