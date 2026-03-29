package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing %d\n", id, j)
		time.Sleep(time.Second)
	}
}

func main() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send jobs
	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
}
