package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go uploadedFile3(&wg)
	wg.Add(1)
	go uploadedFile2(&wg)
	wg.Add(1)
	go uploadedFile1(&wg)

	wg.Wait()

}
func uploadedFile1(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("file1 loaded")
	wg.Done()
}
func uploadedFile2(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("file1 loaded")
	wg.Done()

}
func uploadedFile3(wg *sync.WaitGroup) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("file1 loaded")
	wg.Done()

}
