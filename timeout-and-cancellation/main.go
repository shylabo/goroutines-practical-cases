package main

import (
	"fmt"
	"time"
)

/**
- Timeout and Cancellation -

Channels can be utilized to implement timeouts or cancellations in goroutines
by using a select statement with a channel and a time.
After channel, you can set a timeout for a specific operation.
If the operation doesn't complete within the specified time, the timeout channel
will be selected, and you can handle the timeout accordingly.
Similarly, you can use a cancellation channel to signal goroutines to stop their
execution when a certain condition occurs.
*/

func operation(ch chan<- bool) {
	time.Sleep(3 * time.Second)
	ch <- true // Send true to indicate operation completed
}

func main() {
	ch := make(chan bool)

	go operation(ch)

	select {
	case <-ch:
		fmt.Println("Operation completed.")
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout occurred.")
	}
}
