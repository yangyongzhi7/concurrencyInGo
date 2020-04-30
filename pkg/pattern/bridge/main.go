package main

import "fmt"

func main() {
	done := make(chan interface{})
	defer close(done)

	valStream := Bridge(done, GenerateChannels(100))
	for v := range valStream {
		fmt.Println(v)
	}
}
