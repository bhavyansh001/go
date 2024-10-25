package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
	days := []string{"Sunday", "Tuesday", "Webnesday", "Friday", "Saturday"} // slice it is

	fmt.Println(days)

	// Loops
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])		
	}

	//Better
	for i := range days {
		fmt.Println(days[i])
	}

	// for each
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index + 1, day)
	}

	// if no index:
	// for _, day ...

	rougeValue := 1

	for rougeValue < 10 {
		// if rougeValue == 5 {
		// 	break
		// }
		if rougeValue == 5 {
			rougeValue++
			continue
		}

		if rougeValue == 2 {
			goto diversepixel
		}

		fmt.Println("Value is: ", rougeValue)
		rougeValue++
	}

	diversepixel:
		fmt.Println("Jumping at diversepixel.com")
}