package task26

import (
	"fmt"
	"strings"
)

func CheckUniq(str string) bool {
	// Преобразуем строку к нижнему регистру для регистронезависимой проверки
	str = strings.ToLower(str)
	data := []byte(str)
	// Создаем карту для отслеживания уникальных символов
	mapData := make(map[byte]bool)
	// Перебираем каждый символ в строке
	for _, item := range data {
		// Если символ уже есть в карте, значит он не уникален, возвращаем false
		if mapData[item] {
			return false
		}
		// Добавляем символ в карту
		mapData[item] = true
	}
	// Если прошли через всю строку без повторений, возвращаем true
	return true
}

// CheckUniq2 Если необходимо обрабатывать Unicode
func CheckUniq2(input string) bool {
	// Преобразуем строку к нижнему регистру для регистронезависимой проверки
	input = strings.ToLower(input)

	// Создаем карту для отслеживания уникальных символов
	charMap := make(map[rune]bool)

	// Перебираем каждый символ в строке
	for _, char := range input {
		// Если символ уже есть в карте, значит он не уникален, возвращаем false
		if charMap[char] {
			return false
		}

		// Добавляем символ в карту
		charMap[char] = true
	}

	// Если прошли через всю строку без повторений, возвращаем true
	return true
}

func Task26() {
	line := "AbcDd"
	if CheckUniq(line) {
		fmt.Println("All letters unique")
	} else {
		fmt.Println("Letters repeating")
	}

}
