package main

import "runtime"

// fanOut
func fanOut(done <-chan interface{}, inStream <-chan int, fn func(done <-chan interface{}, in <-chan int) <-chan int) []<-chan int{
	cpuNum := runtime.NumCPU()
	subChannels := make([]<-chan int, cpuNum)

	for i := 0; i < cpuNum; i++ {
		subChannels[i] = fn(done, inStream)
	}

	return subChannels

}
