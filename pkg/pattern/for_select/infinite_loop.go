package main

import (
	"fmt"
	"time"
)

// Looping infinitely waiting to be stopped
func main() {
	done := make(chan interface{})

	go func() {
		fmt.Printf("Start goroutine\n")
		time.Sleep(5 * time.Second)
		close(done)
	}()

	for {
		select {
		case <-done:
			fmt.Printf("Read from done channel\n")
			return
		default:
			// Dome some interesting
			time.Sleep(500 * time.Microsecond)
			fmt.Printf("deault action\n")
		}
	}

	fmt.Printf("Done\n")
}
