package main

import (
	"fmt"
)

func main() {

	defer fmt.Println("World")

	// defer will execute the statement just before the function ends
	defer fmt.Println("One")

	defer fmt.Println("Two")

	fmt.Println("Hello")

	// Output: Hello Two One World
	myDefer()

	// Output: Hello 43210Two One World
}

func myDefer(){
	for i:=0; i<5; i++{
		defer fmt.Print(i)
	}
}