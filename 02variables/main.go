package main

import "fmt"

// jwtToken := 3000000 // this will not work, as walrus operator is only allowed inside functions

// First letter caps means it is public, can be accessed outside the package
const LoginToken string = "asjdhfjhasdfjhasdfjhasdf"


func main() {

	var username string = "Atharva Deshmukh"
	fmt.Println("Hello, ", username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	// 5 decimals in float32
	var smallFloatValue float32 = 255.123456789
	fmt.Println(smallFloatValue)
	fmt.Printf("Variable is of type: %T \n", smallFloatValue)


	// default values and aliases
	// for int it is 0, not garbage value
	var anotherVariable int;
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type, lexer will decide the type
	var website  = "https://atharva-builds.vercel.app"
	fmt.Println(website)

	// no var style, (:) is used for short declaration and called as walrus operator
	// only works inside functions
	numberOfUsers := 300000.001
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken);

	fmt.Printf(("Variable is of type: %T \n"), LoginToken)
}