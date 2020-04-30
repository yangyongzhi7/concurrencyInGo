package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork(done <-chan interface{}, id int, wg *sync.WaitGroup, results chan<- int) {
	started := time.Now()
	defer wg.Done()

	// simulate random load
	simulatedLoadTime := time.Duration((1 + rand.Intn(5))) * time.Second
	select {
	case <-done:
	case <-time.After(simulatedLoadTime):
	}

	select {
	case <-done:
	case results <- id:
	}

	took := time.Since(started)
	//
	if took < simulatedLoadTime {
		took = simulatedLoadTime
	}

	fmt.Printf("%v took %v\n", id, took)

}
