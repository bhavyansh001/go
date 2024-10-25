package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	// PerformPostJsonRequest()
	// PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "https://foodfusion.diversepixel.com/api/v1/restaurants"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is: ", byteCount)

	fmt.Println(responseString.String())

	fmt.Println(string(content)) // just do this avoiding strings.Builder for simplicity, get content and just print it
}

func PerformPostJsonRequest() {
	const myurl = "https://foodfusion.diversepixel.com/api/v1/login"

	requestBody := strings.NewReader(`
		{
			"email":"alice.johnson@example.com",
			"password":"password789"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const someurl = "http://somesite.com:80/endpoint"

	//formdata

	data := url.Values{}

	data.Add("firstname", "Bhavyansh")
	data.Add("lastname", "Yadav")
	data.Add("email", "some@example.com")

	response, _ := http.PostForm(someurl, data)

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}