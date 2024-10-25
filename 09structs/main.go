package main

import "fmt"

func main()  {
	fmt.Println("Structs in golang")

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 34}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh's details are: %+v\n" , hitesh)
	fmt.Printf("Name is: %+v and email is: %v\n" , hitesh.Name, hitesh.Email)
	
}

type User struct { // first letter capital as public
	Name	string // press tab
	Email 	string
	Status 	bool
	Age 	int
}
