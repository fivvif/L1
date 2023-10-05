package task20

import (
	"fmt"
	"strings"
)

func Task20() {
	str := "I want a new hat aboboaboa"
	words := strings.Fields(str)
	newString := ""
	for i := len(words) - 1; i >= 0; i-- {
		newString += words[i] + " "

	}
	fmt.Println(newString)
}
