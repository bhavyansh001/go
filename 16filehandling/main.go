package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file - diversepixel.com"

	file, err := os.Create("./mydpgofile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("Length is: ", length)

	 defer file.Close()  // recommended way

	 readFile("./mydpgofile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)
	// fmt.Println("Text data inside the file is \n", databyte)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

// data is read as bytes, not text


func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}