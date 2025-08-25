package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	timer := time.NewTimer(4 * time.Second)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "данные" //"Долго выполняем операцию"
	}()
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				ch = nil
			} else {
				fmt.Println("Получены", msg)
			}
		case <-timer.C:
			fmt.Println("Подождем еще немного")
		case <-time.After(10 * time.Second):
			fmt.Println("Дождались")
			ch = nil
		}
		if ch == nil {
			break
		}
	}
}
