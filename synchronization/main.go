package main

import (
	"fmt"
	"sync"
)

/**
- Synchronization -

Channels can be used to synchronize the execution of multiple goroutines.
By using channels to coordinate actions, you can ensure that certain operations
are performed in a specific order or that specific conditions are met before proceeding.
For example, you can use a channel to signal when a group of goroutines has completed
their tasks before continuing with the main program flow.
*/

func worker(wg *sync.WaitGroup, task int) {
	defer wg.Done() // Notify WaitGroup when task is done
	fmt.Println("Processing task", task)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment WaitGroup counter for each task
		go worker(&wg, i)
	}

	wg.Wait() // Wait until all tasks are completed
	fmt.Println("All tasks completed")
}
