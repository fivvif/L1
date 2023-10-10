package task11

import "fmt"

func Task11() {
	set1 := []int{1, 6, 4, 2, 0}
	set2 := []int{6, 8, 3, 2, 4}
	// создаём мапу для проверки наличия цифры в множестве
	set1Map := make(map[int]bool)
	for _, item := range set1 {
		// заполняем мапу первым множеством
		set1Map[item] = true
	}

	var result []int

	for _, item := range set2 {
		// идем по второму множеству и если элемент с таким ключом существует(след. существуем в обоих множествах)
		// то добавляем в результат.
		if set1Map[item] {
			result = append(result, item)
		}
	}
	fmt.Printf("Intersection of sets : %d \n", result)

}
