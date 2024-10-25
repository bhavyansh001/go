package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	// Application which whenever opens up gives time on CLI
	// go env
	// GOOS="windows" go build
	// GOOS="darwin" go build
	// GOOS="linux" go build
}

// Memory management
// new(), make() used more, initialized also, gives non-zeroed storage