package task2

import (
	"fmt"
	"sync"
)

func Task2() {
	arr := [5]int{2, 4, 6, 8, 10} // специально массив, а не срез
	var wg sync.WaitGroup
	for _, num := range arr {
		//инкрементим счетчик при создании горутины
		wg.Add(1)
		go func(n int) {
			//декрементим счётчик когда функция завершит работу
			defer wg.Done()
			square := n * n
			fmt.Println(square)
		}(num)
	}
	//ждёт когда счётчик станет равен нулю, тоесть когда все горутины доработают
	wg.Wait()

}
