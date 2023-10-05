package task12

import "fmt"

func Task12() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	for _, value := range strings {
		set[value] = true
	}

	for key := range set {
		fmt.Println(key)
	}
}
