package main

import "fmt"

func main() {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)

		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()

		return intStream
	}

	multiply := func(done <-chan interface{}, inputStream <-chan int, multiplied int) <-chan int {
		outStream := make(chan int)

		go func() {
			defer close(outStream)
			for v := range inputStream {
				select {
				case <-done:
					return
				case outStream <- v * multiplied:
					fmt.Printf("multiplied :%d\n", v*multiplied)
				}
			}
		}()

		return outStream
	}

	add := func(done <-chan interface{}, inputStream <-chan int, additive int) <-chan int {
		outStream := make(chan int)

		go func() {
			defer close(outStream)
			for v := range inputStream {
				select {
				case <-done:
					return
				case outStream <- v + additive:
					fmt.Printf("additive :%d\n", v+additive)
				}
			}
		}()

		return outStream
	}

	done := make(chan interface{})
	defer close(done)

	for result := range add(done, multiply(done, generator(done, 1, 2, 3, 4), 2), 1) {
		fmt.Println(result)
	}
}
