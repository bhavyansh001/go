package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // usually they are pointers

var mut sync.Mutex	// pointer is created usually

// multiple servers: wait groups, like in read-replicas, microservices, multi push, bringing data from different endpoints
// for running and working with goroutines: use package sync

func main() {
	// go greeter("Hello") 	// goroutine: go
	// greeter("world")
	websitelist := []string{
		"https://diversepixel.com",
		"https://go.dev",
		"https://google.co.in",
		"https://github.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()		// at the end of the method being called out
	fmt.Println(signals)
	// Mutex: mutual exclusion lock, provides lock on memory for one goroutine to work at a time
	// R/W Mutex
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 	// 	time.Sleep(3000 * time.Microsecond) // or sync package
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Wrong")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s website", res.StatusCode, endpoint)
	}
}