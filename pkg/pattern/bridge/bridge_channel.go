package main

import (
	"concurrencyInGo/pkg/pattern/stages"
	"context"
)

// bridge
func Bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	context.Background()
	go func() {
		defer close(valStream)

		for {
			var stream <-chan interface{}
			select {
			case <-done:
				return
			case maybeStream, ok := <-chanStream:
				if !ok {
					return
				}
				stream = maybeStream
			}

			for val := range stages.OrDone(done, stream) {
				select {
				case <-done:
				case valStream <- val:
				}
			}
		}
	}()

	return valStream
}
