package task6

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Task6() {
	//1 канал для выхода
	exitCh1 := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			//ждём записи и получая ее выходим из горутины
			case <-exitCh1:
				fmt.Println("Stopping 1")
				return

			default:
				time.Sleep(200 * time.Millisecond)
				fmt.Println("Working 1")
			}

		}
	}()

	time.Sleep(2 * time.Second)
	close(exitCh1)
	wg.Wait()
	time.Sleep(1 * time.Second)
	// 2 контекст
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopping 2")
				return
			default:
				time.Sleep(200 * time.Millisecond)
				fmt.Println("Working 2")
			}

		}
	}()
	wg.Wait()
	// с помощью пакета runtime
	kanal := make(chan string)
	alarm := "Viklu4ai!!!"
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-kanal:
				fmt.Println("Ponyal vilku4au")
				runtime.Goexit()
			default:
				time.Sleep(700 * time.Millisecond)
				fmt.Println("Rabotau...")
			}
		}
	}()
	time.Sleep(1 * time.Second)
	kanal <- alarm
	wg.Wait()
	fmt.Println("1")
}
