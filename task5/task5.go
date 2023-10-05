package task5

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Task5(N int) {
	producer := make(chan int)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	// горутина пишущая в канал producer пока не получить сигнал cancel()
	go func() {
		defer wg.Done()
		defer close(producer)
		for i := 0; ; i++ {
			select {
			case producer <- i:
			case <-ctx.Done():
				fmt.Println("Time expired")
				return
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		// читает пока открыт канал, а он закроется после завершения работы писателя
		for info := range producer {
			fmt.Println(info)
		}

	}()
	// main ждёт 5 секунд
	time.Sleep(time.Duration(N) * time.Second)
	// отправляет сигнал завершения
	cancel()
	//ждём всех
	wg.Wait()
	fmt.Println("Finish")
}
