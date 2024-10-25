package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5  // add 5 in myCh
	// fmt.Println(<-myCh)
	wg.Add(2)
	// <-chan: read only channel
	go func(ch <-chan int, wg *sync.WaitGroup){
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)


		fmt.Println(<-myCh)
		// fmt.Println(<-myCh) // to avoid this, use buffered channel, `2` in myCh
		wg.Done()
	}(myCh, wg)

	// send only channel
	go func(ch chan<- int, wg *sync.WaitGroup){
		// myCh <- 5
		myCh <- 0 // also returns zero when no sender provided, hence used isChannelOpen
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
// Channels are a way for multiple goroutines (threads in golang) to talk to each other

// error: deadlock, if there is a value passed, there should be a listener listening to that value