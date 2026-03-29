package main

import (
	"fmt"
	"time"
)

func worker(worker int, jobch <-chan int, resultch chan<- int, done chan<- bool) {
	for ch := range jobch {
		fmt.Printf("worker %d received job %d \n", worker, ch)
		time.Sleep(2 * time.Second)
		resultch <- ch
	}
	done <- true // signal this worker is finished
}

func main() {
	const (
		jobs  = 10
		workr = 3
	)

	jobch := make(chan int, jobs)
	resultch := make(chan int, jobs)
	done := make(chan bool, workr)

	// Start workers
	for i := 1; i <= workr; i++ {
		go worker(i, jobch, resultch, done)
	}

	// Send jobs
	for j := 1; j <= jobs; j++ {
		jobch <- j
	}
	close(jobch)

	// Close resultch AFTER all workers are done
	go func() {
		for i := 0; i < workr; i++ {
			<-done // wait for each worker
		}
		close(resultch)
	}()

	// Read results safely
	for rch := range resultch {
		fmt.Println("Result:", rch)
	}
}
