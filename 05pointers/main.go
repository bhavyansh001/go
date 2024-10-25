package main
// pointers: surety that passed on thing is actual value from memory and not a copy. Used when we want to pass actual value.

import "fmt"

func main()  {
	fmt.Println("Welcome to class on pointers")

	// var ptr *int // pointer initialization, nil by default

	// fmt.Println("Value of ptr is ", ptr)

	myNumber := 23

	var ptr = &myNumber  // creating a pointer when referencing a value, a reference to direct memory location

	fmt.Println("Value of actual pointer is ", ptr) // memory address
	fmt.Println("Value of actual pointer is ", *ptr) // the value

	*ptr = *ptr * 2
	fmt.Println("New value is: ", myNumber)
	
}