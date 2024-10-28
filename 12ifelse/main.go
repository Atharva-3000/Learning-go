package main

import "fmt"

func main() {
	fmt.Println("If else statement in Golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 && loginCount < 100 {
		result = "Special User"
	} else {
		result = "Super User"
	}

	fmt.Println("User is a: ", result)

	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	// if with short statement and at point declaration
	if num := 3; num < 10 {
		fmt.Println(num, "is less than 10")
	} else {
		fmt.Println(num, "is not less than 10")
	}

	// error handling
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
