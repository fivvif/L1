package task11

import "fmt"

func Task11() {
	set1 := []int{1, 6, 4, 2, 0}
	set2 := []int{6, 8, 3, 2, 4}
	set1Map := make(map[int]bool)
	for _, item := range set1 {
		set1Map[item] = true
	}

	var result []int

	for _, item := range set2 {
		if set1Map[item] {
			result = append(result, item)
		}
	}
	fmt.Printf("Intersection of sets : %d \n", result)

}
