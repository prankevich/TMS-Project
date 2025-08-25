package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			counter++
			mu.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			counter++
			mu.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Counter:", counter)
}
