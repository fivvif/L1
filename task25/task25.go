package task25

import (
	"fmt"
	"time"
)

// Sleep используем метод After
func Sleep(delay time.Duration) {
	<-time.After(delay)
}

// другим методом
func Sleep2(delay time.Duration) {
	// начинаем отсчёт времени
	start := time.Now()
	for {
		// проверяем счётчик, пока он не станет равным необходимому времени
		if time.Since(start) == delay {
			break
		}
	}
}

func Task25() {
	//функция sleep выглядит короче и лаконичнее, чем sleep2
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
