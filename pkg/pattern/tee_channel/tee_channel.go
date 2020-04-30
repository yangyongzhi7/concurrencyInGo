package main

import "concurrencyInGo/pkg/pattern/stages"

// tee
func tee(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})

	go func() {
		defer close(out1)
		defer close(out2)

		for val := range stages.OrDone(done, in) {
			var o1, o2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
				case o1 <- val:
					o1 = nil
				case o2 <- val:
					o2 = nil
				}
			}
		}
	}()

	return out1, out2
}
