package main

import (
	"fmt"
	"time"
)

func load(id int, ch chan string) {
	time.Sleep(time.Duration(id) * time.Second)
	ch <- fmt.Sprintf("Горутина %d завершила загрузку", id)
}

func main() {
	ch := make(chan string)

	for i := 1; i <= 5; i++ {
		go load(i, ch)
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(<-ch)
	}
}
