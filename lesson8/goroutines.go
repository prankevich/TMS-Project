package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go evenNumbers(&wg)
	wg.Add(1)
	go oddNumbers(&wg)
	wg.Wait()
}
func evenNumbers(wg *sync.WaitGroup) {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Cписок четных чисел:", i)
		}
	}
	wg.Done()

}
func oddNumbers(wg *sync.WaitGroup) {
	for i := 1; i <= 9; i++ {
		if i%2 != 0 {
			fmt.Println("Cписок нечетных чисел:", i)
		}
	}
	wg.Done()
}
