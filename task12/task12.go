package task12

import "fmt"

func Task12() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)
	// каждый элемент в множестве уникальный, поэтому удобнее всего использовать карту для хранения
	for _, value := range strings {
		set[value] = true
	}
	fmt.Println(set)
	for key := range set {
		fmt.Println(key)
	}
}
