package main

import (
	"fmt"
	"sync"
)

func main() {
	done := make(chan interface{})
	results := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go doWork(done, i, &wg, results)
	}

	firstReturned := <-results
	close(done)
	wg.Wait()

	fmt.Printf("Received an answer from #%v\n", firstReturned)

}
