// variables

package main

import "fmt"

const LoginToken string = "someval"  // 'L': public constant declaration

func main() {
  var username string = "bhavyansh"
  fmt.Println(username)  // short: fp
  fmt.Printf("Variable is of type: %T \n", username)

  var isLoggedIn bool = false
  fmt.Println(isLoggedIn)
  fmt.Printf("Variable is of type: %T \n", isLoggedIn)

  var smallVal uint8 = 255 // (largest possible)
  fmt.Println(smallVal)
  fmt.Printf("Variable is of type: %T \n", smallVal)

  var smallFloat float32 = 255.45523432456901  // float64 will get more precise data
  fmt.Println(smallFloat)
  fmt.Printf("Variable is of type: %T \n", smallFloat)

  // default values and aliases
  var anotherVar int  //default value that goes in is 0, no garbage value
  fmt.Println(anotherVar)
  fmt.Printf("Variable is of type: %T \n", anotherVar)

  // Implicit type of variable declaration
  var website = "diversepixel.com"
  fmt.Println(website)
  fmt.Printf("Variable is of type: %T \n", website)

  // no var style
  numberOfUser := 300000.0  // walrus operator is allowed in functions, not outside them
  fmt.Println(numberOfUser)

  fmt.Println(LoginToken)
}
