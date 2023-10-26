package task17

import (
	"fmt"
)

func Task17() {
	// также можно использовать sort.Search(), но опять же в моём понимании суть другая.
	array := []int{1, 2, 4, 6, 10, 12, 15, 17, 20}
	target := 6
	index := binarySearch(array, target)

	if array[index] == target {
		fmt.Printf("Target : %d, index : %d \n", array[index], index)
	}
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid // Элемент найден, возвращаем его индекс
		} else if arr[mid] < target {
			left = mid + 1 // Искать справа от mid
		} else {
			right = mid - 1 // Искать слева от mid
		}
	}

	return -1 // Элемент не найден
}
