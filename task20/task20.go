package task20

import (
	"fmt"
	"strings"
)

func Task20() {
	str := "I want a new hat aboboaboa"
	// создаем слайс с каждым словом(разделяем строку по пробелам)
	words := strings.Fields(str)
	newString := ""
	// записываем каждое слово в новую строку, идя с конца массива и дописывая пробел
	for i := len(words) - 1; i >= 0; i-- {
		newString += words[i] + " "

	}
	fmt.Println(newString)
}
