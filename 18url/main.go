package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://diversepixel.com:443/portfolio/chatapp.html"

func main() {
	fmt.Println("Welcome to handling urls")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery) // query params

	qparams := result.Query() // all query params in a better fomat
	fmt.Printf("The type of query params are: %T\n", qparams) // url.Values = key value pairs

	fmt.Println(qparams["key"]) // returns the value

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	// url from params
	partsOfUrl := &url.URL{ // reference is passed
		Scheme: "https",
		Host: "diversepixel.com",
		Path: "/portfolio",
		RawPath: "user=bhavyansh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}