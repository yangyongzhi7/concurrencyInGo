package main

import "time"

// a goroutine which can be managed by a steward
type startGoroutineFn func(done <-chan interface{}, pulseInterval time.Duration) (heartbeat <-chan interface{})


 a := func(timeout time.Duration, startGoroutine startGoroutineFn) startGoroutineFn{
	return func(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}) {
heartbeat:=make(chan interface{})
return heartbeat
go func() {
	defer close(heartbeat)


}()


	}
	
}


