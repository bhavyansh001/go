package main

// pkg.go.dev
// bufio
// os

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// welcome := "Welcome to user input"
	// fmt.Println(welcome)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter your rating: ")

	// // comma ok || comma error syntax; it doesn't have try catch for error handling
	// input, _ := reader.ReadString('\n')  // _ if error is not concern, ideal is , err
	// fmt.Println("Thanks for rating ", input)
	// fmt.Printf("Type of this rating is %T", input)

	// Num input
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our pizza")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating + 1)
	}
}