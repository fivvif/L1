package task4

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

//Та же реализиция только  через контекс

func worker(ctx context.Context, wg *sync.WaitGroup, ch <-chan int, id int) {
	defer wg.Done()
	for {
		select {
		//беспрерывно читает из канала ch, в случае если чтение не удалось worker завершается
		case info, ok := <-ch:
			if !ok {
				fmt.Printf("Main channel close, worker %d stopping \n", id)
				return
			}
			fmt.Printf("Recived message : %d \n", info)
		//ждём сигнала от контекста и при его получее завершаем воркера
		case <-ctx.Done():
			fmt.Printf("Worker %d finish working \n", id)
			return
		}
	}
}

func Task4_2(N int) {
	mainChannel := make(chan int)
	// создаём контекст с отменой
	mainCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	// создаём воркеров
	for counter := 0; counter <= N; counter++ {
		wg.Add(1)
		go worker(mainCtx, &wg, mainChannel, counter)
	}
	// та же логика что в первой реализации
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	wg.Add(1)
	go func() {
		defer close(mainChannel)
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case mainChannel <- i:
			// получая сигнал quit отменяем контекс
			case <-quit:
				fmt.Println("Quit signal")
				cancel()
				return
			}

		}
	}()

	wg.Wait()
	fmt.Println("Finish")

}
