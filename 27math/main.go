package main

import (
	// "math/rand"
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Maths in Golang")

	// var myNum int = 2
	// var Num2 float64 = 4

	// fmt.Println("The sum is: ", myNum+int(Num2))

	// random number
	// fmt.Println(rand.Intn(5))

	// random from crypto
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}