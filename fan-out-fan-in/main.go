package main

import (
	"fmt"
	"sync"
)

/**
- Fan-Out/Fan-In -

Channels can be used to distribute work across multiple goroutines (fan-out)
and then collect the results from all goroutines (fan-in). This approach is useful
when you have a large amount of work to be done, and splitting it among multiple goroutines
can speed up the overall processing time. Each goroutine can perform its computation and
send the results through a channel. The main goroutine can then receive the results
from all the channels and aggregate them.
*/

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		// Perform job processing here
		results <- job * 2 // Send result too results channel
	}
}

func main() {
	numJobs := 10
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Spawn worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			worker(workerID, jobs, results)
			wg.Done()
		}(i)
	}

	// Send jobs to jobs channel
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs) // Close jobs channel to indicate no more jobs

	wg.Wait() // Wait for all workers to complete

	// Collect results from results channel
	for i := 1; i <= numJobs; i++ {
		fmt.Println(<-results)
	}
}
