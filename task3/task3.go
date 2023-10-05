package task3

import (
	"fmt"
	"sync"
)

func Task3() {
	arr := []int{2, 4, 6, 8, 10}
	// та же логика с wg
	var wg sync.WaitGroup
	ch := make(chan int)
	for _, num := range arr {
		wg.Add(1)
		//записываем квадраты в канал
		go func(n int) {
			defer wg.Done()
			square := n * n
			ch <- square
		}(num)
	}
	//горутина закрывающая канал, после того как все горутины доработают и отправят все данные
	go func() {
		wg.Wait()
		close(ch)
	}()
	sum := 0
	//цикл читающий из канала ch, пока он не будет закрыт
	for square := range ch {
		sum += square
	}

	fmt.Println(sum)

}
