package main

import (
	"fmt"
	"math/rand"
)

func useBufferedChannels(n int) {

	// un-buffered channels take only one piece of content, then block to wait for a read or write
	// Buffered channels on the other hand, take n values before blocking
	// Un-buffered channels are predictable, because we know when the goroutine will block

	// Buffered channels are useful when
	// 1. We are running multiple goroutines and we want to block until all the routines finish
	// and
	// 2. We know ahead of time, the number of routines we want
	//

	ch3 := make(chan int, n) // this open buffer will need 10 elements to unblock

	for i := 0; i < n; i++ {
		// Here we launch 10 go routines each writing to the channel
		// any less or any more, we'd block either one of the go routines, because we'd read more values than would have been
		// written or the vice-versa thus one of the go routines would block causing a deadlock
		go func() {
			v := rand.Int()
			fmt.Printf("Writing %dth value %d\n", i+1, v)
			ch3 <- v
		}()
	}

	fmt.Println("Reading from /\\channel 3/\\")

	for i := 0; i < n; i++ {
		randoms := <-ch3
		fmt.Printf("Reading %dth value %d\n", i+1, randoms)
	}
	close(ch3)
}

// Go's concurrency is based on Communicating Sequential Processes
func main() {
	ch := make(chan string)  // We've made a channel
	ch2 := make(chan string) // another channel
	go func() {
		// Within this goroutine, we loop a select
		// The select statement below looks for available next steps.
		// It checks if channel 1 is available for reading from, if not, it doesn't block
		// Instead it drops down to check if channel 2 is available for writing
		// Because we started execution on the main thread, it sort of ochestrates this back and forth communication
		// Without this for-select, the outer goroutine would need to be very order specific to avoid deadlocks
		for {
			select {
			case msg := <-ch:
				fmt.Println("Message from /\\Channel 1/\\", msg)
			case ch2 <- "Hello Traveler":
				fmt.Println("Message written into /\\Channel 2/\\")
				return
			}
		}
	}()

	// We write into a channel. Remember the main function runs itself as a goroutine
	// So we write into channel 1. This pauses the main goroutine, until another goroutine reads from channel 1
	// So the goroutine above executes
	fmt.Println("Writing into /\\channel 1/\\")
	ch <- "Send me something"

	// the main goroutine resumes and then reads from ch2
	// The main goroutine then pauses from here until another routine writes into ch2
	msg := <-ch2
	fmt.Println("Message from /\\channel 2/\\", msg)

	useBufferedChannels(5)

	fmt.Println("Closing channels")
	close(ch)
	close(ch2)

}
