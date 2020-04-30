package main

import (
	"concurrencyInGo/pkg/pattern/stages"
	"fmt"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	inStream := stages.OrDone(done, stages.Take(done, stages.RepeatFn(done, stages.RandomInt), 100))
	out1, out2 := tee(done, inStream)

	for val := range out1 {
		fmt.Printf("out1: %v,out2: %v\n", val, <-out2)
	}

}
