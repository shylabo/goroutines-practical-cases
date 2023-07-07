package main

import "fmt"

/**
- Parallel Processing -

Channels can be used to distribute work across multiple goroutines for parallel processing.
You can create a channel and spawn multiple goroutines to perform the same task on different portion of the input data.
Each goroutine can process its portion independently and send the results back through the channel.
The main goroutine can then receive and aggregate the results.
*/

func process(data int, result chan<- int) {
	result <- data * 2 // Process data and send result to channel
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	result := make(chan int)

	for _, val := range data {
		go process(val, result) // Spawn goroutines to process data
	}

	for i := 0; i < len(data); i++ {
		fmt.Println(<-result) // Receive and print processed results
	}
}
