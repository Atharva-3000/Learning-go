package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input program"
	fmt.Println(welcome)

	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for my GitHub: ");

	// comma ok aka error ok syntax
	input, _ := reader.ReadString('\n');
	fmt.Println("Thanks for rating me:",input);
	fmt.Printf("Type for rating is %T",input);
}