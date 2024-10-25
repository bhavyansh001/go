package main

import "fmt"

func main(){
	fmt.Println("Welcome to functions in golang")
	// greeter // reference of the function
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes := proAdder(2, 5, 8, 7, -2)
	fmt.Println("Proresult is: ", proRes)

	proRes2, myMessage := proAdder2(2, 5, 8, 7, -2)
	fmt.Println("Proresult is: ", proRes2)
	fmt.Println("Proresult message is: ", myMessage)

}	// main acts as an entrypoint for your program, main function automatically executes

func adder(valOne int, valTwo int) int {		// functions signature: what kind of value you expect to return back
	return valOne + valTwo
}

func proAdder(values ...int) int{	// values is now a slice, ... is variadic functions, which can expect any value
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func proAdder2(values ...int) (int, string) {	// values is now a slice, ... is variadic functions, which can expect any value
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi Pro result function"
}

func greeter() {
	fmt.Println("Namaste from golang")
}

// func () {}
// immediately invoked functions exist (IIFs), anonymous function like the above one exists