package task9

import (
	"fmt"
	"sync"
)

func Task9() {
	inputChannel := make(chan int)
	outputChannel := make(chan int)
	wg := sync.WaitGroup{}
	slice := []int{1, 2, 3, 4, 5}

	wg.Add(1)
	go writeToChannel(&wg, slice, inputChannel)

	wg.Add(1)
	go doubleAndWrite(&wg, inputChannel, outputChannel)

	wg.Add(1)
	go readFromOutput(&wg, outputChannel)

	wg.Wait()
}

func writeToChannel(wg *sync.WaitGroup, arr []int, inputChannel chan<- int) {
	defer wg.Done()
	for _, item := range arr {
		inputChannel <- item
	}

	close(inputChannel)
}

func doubleAndWrite(wg *sync.WaitGroup, inputChannel <-chan int, outputChannel chan<- int) {
	defer wg.Done()

	for num := range inputChannel {
		num *= 2
		outputChannel <- num
	}

	close(outputChannel)
}

func readFromOutput(wg *sync.WaitGroup, outputChannel <-chan int) {
	defer wg.Done()
	for out := range outputChannel {
		fmt.Println(out)
	}
}
