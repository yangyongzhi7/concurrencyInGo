package main

import (
	"fmt"
	"math/rand"
)

// we stipulate a convention:
// If a goroutine is responsible for creating a new goroutine,
// it is also responsible to ensuring it can stop the goroutine.
func main() {
	done := make(chan interface{})
	newRandomStream := func(done <-chan interface{}) <-chan int {
		randomStream := make(chan int)
		go func() {
			for {
				select {
				case randomStream <- rand.Int():
					fmt.Println("Write an integer onto the channel.")
				case <-done:
					return
				}
			}
		}()
		return randomStream
	}

	randomStream := newRandomStream(done)
	for i := 0; i < 3; i++ {
		fmt.Println(<-randomStream)
	}

	close(done)
	fmt.Println("Done")

}
