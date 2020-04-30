package main

import (
	"fmt"
	"sync"
)

var count int
var lock sync.Mutex

func main() {
	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Decrementing: %d\n", count)
	}
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Printf("Arithmetic complete.")

}
