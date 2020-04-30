package main

import "math/big"

func primeFinder(done <-chan interface{}, valueStream <-chan int) <-chan int {
	primeStream := make(chan int)

	go func() {
		defer close(primeStream)

		for v := range valueStream {
			select {
			case <-done:
				return
			default:
				i := big.NewInt(int64(v))
				if i.ProbablyPrime(1) {
					primeStream <- v
				} else {
					primeStream <- 0
				}
			}
		}
	}()

	return primeStream

}
