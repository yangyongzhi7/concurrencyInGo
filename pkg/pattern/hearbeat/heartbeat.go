package main

import (
	"fmt"
	"time"
)

func main() {
	const timeout = 2 * time.Second

	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() { close(done) })

	heartbeat, results := doWork(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Printf("pulse\n")
		case r, ok := <-results:
			if ok == false {
				return
			}
			fmt.Printf("result %v\n", r.Second())
		case <-time.After(timeout):
			fmt.Printf("time after")
			return
		}
	}

}
