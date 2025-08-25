package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	chBuf := make(chan int, 3)
	startTimer := time.Now()

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		chBuf <- 1
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch <- 2
		chBuf <- 2
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 3
		chBuf <- 3
	}()

	fmt.Println("Небуферизированный канал :")
	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Printf("%s Получено: %d\n", time.Since(startTimer), msg)
	}
	fmt.Println("Буферизированного канала :")
	for i := 0; i < 3; i++ {
		msgBuf := <-chBuf
		fmt.Printf("%s  Получено: %d\n", time.Since(startTimer), msgBuf)
	}
}
