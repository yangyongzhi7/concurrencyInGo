package main

import (
	"fmt"
)

// Sending iteration variables out on a channel
func main_() {
	done := make(chan interface{}, 1)
	stringStream := make(chan int, 100) // if you specify this less than 100. a deadlock will happen.
	fmt.Println("Start")

	go func() {
		done <- 1
		fmt.Println("Insert onto done")
	}()

	for i := 0; i <= 100; i++ {
		select {
		case stringStream <- i:
			//time.Sleep(1 * time.Second)
			fmt.Println("Insert onto data stream")
		case <-done:
			fmt.Println("Read from done channel")
			return
		}
	}

	fmt.Println("Done")

}
