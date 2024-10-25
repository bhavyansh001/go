package main

import (
	"fmt"
	"sort"
)

//slices are used over arrays in go
//Automatic memory addition for new entries

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apple", "Orange", "Peach"}

	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println("The updated fruitlist is: ", fruitList)

	//slice up your slice
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3])  // last range is non inclusive, will give [Orange Peach]
	fmt.Println(fruitList)
	fruitList = append(fruitList[:3]) // till 3 non inclusive
	fmt.Println(fruitList)

	// Menory management
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 // Out of bound error

	highScores = append(highScores, 555, 666, 321) // works, memory is reallocated

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	// remove value from slice based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}

	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}