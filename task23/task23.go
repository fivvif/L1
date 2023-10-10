package task23

import "fmt"

func Task23() {
	slice := []int{1, 9, 10, 13, 19, 50, 143, 14314}
	// индекс элемента, который мы хотим удалить
	index := 4
	removedElem := slice[index]
	// полностью перезаписываем слайс, добавляя к срезу до нужного индекса срез после нужного индекса
	// используем ... для распаковки слайса
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Printf("Element at index %d with value %d removed\n", index, removedElem)
	fmt.Println(slice)
}
