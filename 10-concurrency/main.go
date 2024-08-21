package main

import "fmt"


// Go's concurrency is based on Communicating Sequential Processes
func main() {
	ch := make(chan string) // We've made a channel
	ch2 := make(chan string) // another channel

	// Go channels are unbuffered by default
	// Every write to an open unbuffered channel, causes the writing goroutine to pause
	// until another goroutine reads from the channel
	// Likewise, a read from an open channel, causes the reading goroutine to pause until another 
	// goroutine writes to the channel


	// Go also has buffered channels that take certain number of writes until it blocks
	// This is how we specify a buffered channel


	// this is a go routine that will carry on parallel to the main function which is launched as a go routine
	go func () {
		// Within this goroutine, we loop a select
		// The select statement below looks for available next steps. 
		// It checks if channel 1 is available for reading from, if not, it doesn't block
		// Instead it drops down to check if channel 2 is available for writing
		// Because we started execution on the main thread, it sort of ochestrates this back and forth communication
		// Without this for-select, the outer goroutine would need to be very order specific to avoid deadlocks
		for {
			select {
				case msg := <- ch:
					fmt.Println("Message from Channel 1", msg)
				case ch2 <- "Hello Traveler":
					fmt.Println("Message written into channel 2")
					return
			}
		}
	} ()

	// We write into a channel. Remember the main function runs itself as a goroutine
	// So we write into channel 1. This pauses the main goroutine, until a goroutine reads from channel 1
	// So the goroutine above executes
	fmt.Println("Writing into channel 1")
	ch <- "Send me something"

	msg := <- ch2
	fmt.Println("Message from ch2", msg)
	
	
	for i := 10; i >= 0; i-- {
		fmt.Println("Counting the other way", i)
	}
}