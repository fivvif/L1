package task18

import (
	"fmt"
	"sync"
)

func Task18() {
	var wg sync.WaitGroup
	counter := Counter{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				counter.Increment()
			}
			fmt.Println(counter.Show())
		}()
	}
	wg.Wait()
	fmt.Println("Done :", counter.count)
}

// Counter используем мютекс чтобы не было гонки данных
type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Show() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count

}
