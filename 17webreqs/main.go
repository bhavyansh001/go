package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://diversepixel.com"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response if of type: %T\n", response)

	// returns a pointer *http.Response, you're not getting a copy

	defer response.Body.Close() // close must

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}