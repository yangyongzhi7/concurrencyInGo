package stages

// orDone
func OrDone(done <-chan interface{}, inStream <-chan interface{}) <-chan interface{} {
	outStream := make(chan interface{})

	go func() {
		defer close(outStream)

		for {
			select {
			case <-done:
				return
			case v, ok := <-inStream:
				if !ok {
					return
				}

				select {
				case outStream <- v:
				case <-done:
				}
			}
		}

	}()

	return outStream
}
