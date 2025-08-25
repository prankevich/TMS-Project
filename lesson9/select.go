package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 1

	}()
	go func() {
		ch2 <- 2

	}()

	for i := 0; i < 100; i++ {
		select {
		case msq1, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				fmt.Println(msq1)
			}
		case msq2, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				fmt.Println(msq2)
			}
		case <-time.After(1 * time.Second):
			ch1 = nil
			ch2 = nil
		}
		if ch1 == nil && ch2 == nil {
			break
		}
	}
}
