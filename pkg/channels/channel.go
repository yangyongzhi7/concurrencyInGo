package main

import "fmt"

func main() {

	dataStream := make(chan int)

	go func() {
		dataStream <- 1
	}()

	fmt.Printf("The data comes from a channel:%d", <-dataStream)

}
