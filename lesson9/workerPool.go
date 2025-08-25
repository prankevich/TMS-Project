package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		result := j * 2
		results <- fmt.Sprintf("Воркер %d получил результат умножения %d  на 2  равный %d", id, j, result)
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan string, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}
}
