package main

import (
	"fmt"
	"sync"
)

func work(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		result := j * 2
		results <- fmt.Sprintf("Воркер %d получил результат умножения %d  на 2  равный %d", id, j, result)
	}

}
 
func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 5)
	results := make(chan string, 5)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go work(w, jobs, results, &wg)
	}
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()

	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}
}
