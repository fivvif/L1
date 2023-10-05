package task25

import (
	"fmt"
	"time"
)

func Sleep(delay time.Duration) {
	<-time.After(delay)
}

func Sleep2(delay time.Duration) {
	start := time.Now()
	for {
		if time.Since(start) == delay {
			break
		}
	}
}

func Task25() {
	fmt.Println("Starting")
	start := time.Now()
	go func() {
		for {
			fmt.Println("I'm writing while main sleeping")
			Sleep(400 * time.Millisecond)
		}

	}()
	Sleep(5 * time.Second)
	fmt.Println("End")
	fmt.Println("Program worked : ", time.Since(start))
}
