package main

import "fmt"

func main() {
	fmt.Println("Functions in GO")
	greeter()
	result := adder(3, 5)
	fmt.Println("Result is:", result)
	proResult, message := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Pro Result is:", proResult)
	fmt.Println("Message is:", message)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Success"
}

func greeter() {
	fmt.Println("Hello Traveller!")
}
