package task4

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func Task4(N int) {
	mainChannel := make(chan int)
	var wg sync.WaitGroup
	//создаем воркеров
	for counter := 0; counter < N; counter++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//читаем из mainChannel пока он не закроется
			for info := range mainChannel {
				fmt.Printf("Recived message : %d \n", info)
			}
			fmt.Println("Worker finished work")
		}()
	}
	//канал для принятия нажатия Ctrl + C
	quit := make(chan os.Signal, 1)
	// ожидает сигнала ctrl c
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	// горутина для отправки сообщений
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(mainChannel)
		for i := 0; ; i++ {
			select {
			//пишет в канал
			case mainChannel <- i:
				//ожидает сигнала quit и заканчивает функцию
			case <-quit:
				return
			}

		}
	}()

	//чтобы все горутины корректно завершили работу
	wg.Wait()
	fmt.Println("Finish")

}
