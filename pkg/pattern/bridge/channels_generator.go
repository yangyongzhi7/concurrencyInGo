package main

func GenerateChannels(count int) <-chan <-chan interface{} {
	chanStream := make(chan (<-chan interface{}))

	go func() {
		defer close(chanStream)
		for i := 0; i < count; i++ {
			stream := make(chan interface{}, 1)
			stream <- i
			close(stream)
			chanStream <- stream
		}

	}()

	return chanStream

}
