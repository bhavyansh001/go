package main

import "fmt"

func main() {
	// fmt.Println("Hello")
	// defer fmt.Println("World")

	defer fmt.Println("Hello")  // defer cuts this line from here and places it just before the curly braces end (///there), as LIFO if multiple present
	fmt.Println("World")

	defer fmt.Println("Two")
	fmt.Println("World")

	myDefer()

///here
}

func myDefer()  {
	for i := 0; i < 5; i++ {
		 defer fmt.Print(i)
	}
}