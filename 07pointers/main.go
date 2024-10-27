package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go class")

	// var ptr *int
	// fmt.Println("Value of ptr is: ", ptr) //prints nil

	// fmt.Println("Address of ptr is: ", &ptr) //prints address of ptr

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Address of myNumber is: ", ptr)
	fmt.Println("Actual value of myNumber is: ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value of myNumber is: ", myNumber)
}
