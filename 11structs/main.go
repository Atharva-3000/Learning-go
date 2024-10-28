package main

import "fmt"

func main() {
	fmt.Println(" Structs in Go")

	// no inheritance in Go
	Atharva := User{"Atharva",
		"atharva@gmail.com",
		true,
		21}

	fmt.Println(Atharva)
	fmt.Printf("Atharva's details are %+v\n", Atharva)

	fmt.Printf("Name is %v and Email is %v\n", Atharva.Name, Atharva.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
