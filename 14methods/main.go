package main

import "fmt"

func main()  {
	fmt.Println("Structs in golang")

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 34}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh's details are: %+v\n" , hitesh)
	fmt.Printf("Name is: %+v and email is: %v\n" , hitesh.Name, hitesh.Email)

	hitesh.GetStatus()
	hitesh.NewMail() // does not change the object, hence a copy of that object is passed along, if want to change, pass reference using pointers

	fmt.Printf("Name is: %+v and email is: %v\n" , hitesh.Name, hitesh.Email)
	
}

type User struct { // first letter capital as public
	Name	string // press tab
	Email 	string
	Status 	bool
	Age 	int
	// oneAge	int	// private, unexportable
}

// method = funcition in a class, in golang structs
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
