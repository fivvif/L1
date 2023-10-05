package task7

import (
	"fmt"
	"sync"
)

func Task7() {
	karta := make(map[string]int)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			mutex.Lock()
			karta[key] = i * 10
			mutex.Unlock()
			fmt.Printf("Горутина %d записала: %s -> %d\n", i, key, i*10)
		}(i)
	}
	var mapka sync.Map
	for i := 100; i < 125; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			mapka.Store(key, i*10)
			fmt.Printf("Горутина %d записала: %s -> %d\n", i, key, i*10)
		}(i)
	}
	wg.Wait()
	for key, value := range karta {
		fmt.Printf("%s -> %d\n", key, value)
	}
	fmt.Println("/////////////////////")
	mapka.Range(func(key, value any) bool {
		fmt.Printf("%s -> %d\n", key, value)
		return true
	})

	fmt.Println(mapka)

}
