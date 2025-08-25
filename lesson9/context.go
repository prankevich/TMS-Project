package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan string)
	for i := 1; i < 10; i++ {
		go loadF(ctx, i, ch)
	}
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()
	for i := 1; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func loadF(ctx context.Context, id int, ch chan string) {
	select {
	case <-time.After(time.Duration(id) * time.Second):
		ch <- fmt.Sprintln("Горутина завершила загрузку ", id)
	case <-ctx.Done():
		ch <- fmt.Sprintln("Горутина отменена", id)
	}
}
