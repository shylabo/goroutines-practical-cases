package main

import "fmt"

/**
- Producer-Consumer Pattern -

Channels are commonly used to implement the producer-consumer pattern,
where one goroutine (producer) generates data and sends it to another
goroutine (consumer) for processing. The producer can send data through a channel,
and the consumer can receive and process it. This pattern is useful for handling
concurrent data processing, such as reading from a file and performing calculations on the data.
*/

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i // Send value to channel
	}
	close(ch) // Close the channel when done
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println(num) // Process value received from channel
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
