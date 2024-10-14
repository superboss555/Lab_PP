package main

import (
	f "fmt"
	s "sync"
)

var counter int = 0

func main() {
	channel := make(chan bool)
	var mutex s.Mutex
	for i := 1; i < 10; i++ {
		go increment(i, channel, &mutex)
	}

	for i := 1; i < 10; i++ {
		<-channel
	}

	f.Println("End")
}

func increment(number int, ch chan bool, mutex *s.Mutex) {
	//mutex.Lock()
	for k := 1; k <= 10; k++ {
		counter++
		f.Println("Goroutine ", number, "-", counter)
	}

	//mutex.Unlock()
	ch <- true
}
