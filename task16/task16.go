package task16

import "fmt"

func Task16() {
	// можно использовать пакет sort, но как я понял суть задания в создании своей функции
	array := []int{2, 1, 6, 7, 4, 3, 10, 111, 41, 55, 78, 1, 0, 4}
	fmt.Println(quickSort(array))
}

func quickSort(array []int) []int {
	// проверяем массив чтобы не уйти в беск рекурсию
	if len(array) <= 1 {
		return array
	}
	// определяем начальный элемент, я сделал его последним в массиве
	pivot := array[len(array)-1]
	var leftArray, rightArray, equals []int
	// идём по исходному массиву и сортируем элементы в массивы  меньшие, большие и равные pivot
	for _, item := range array {
		switch {
		case item < pivot:
			leftArray = append(leftArray, item)
		case item == pivot:
			equals = append(equals, item)
		case item > pivot:
			rightArray = append(rightArray, item)
		}
	}
	//используем рекурсию для сортировки подмассивов
	leftArray = quickSort(leftArray)
	rightArray = quickSort(rightArray)
	// заполняем результат и возвращаем его
	result := append(leftArray, equals...)
	result = append(result, rightArray...)

	return result
}
