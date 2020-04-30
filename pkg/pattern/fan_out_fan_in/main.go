package main

import (
	"concurrencyInGo/pkg/pattern/stages"
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	experimentalCount := 1000000

	// Sequence search
	start := time.Now()
	randIntStream := primeFinder(done, toInt(done, stages.Take(done, stages.RepeatFn(done, stages.RandomInt), experimentalCount)))
	for range randIntStream {
		//fmt.Print(r)
	}
	fmt.Printf("\nsearch took: %v\n", time.Since(start))

	// concurrently search
	start = time.Now()
	randomIntStream := toInt(done, stages.Take(done, stages.RepeatFn(done, stages.RandomInt), experimentalCount))
	finders := fanOut(done, randomIntStream, primeFinder)
	for range fatIn(done, finders...) {
		//fmt.Print(prime)
	}

	fmt.Printf("\nsearch took: %v\n", time.Since(start))

}
