package main

import "fmt"

func main() {
	fmt.Println(" Methods in Go")

	// no inheritance in Go
	Atharva := User{"Atharva",
		"atharva@gmail.com",
		true,
		21}

	fmt.Println(Atharva)
	fmt.Printf("Atharva's details are %+v\n", Atharva)

	fmt.Printf("Name is %v and Email is %v\n", Atharva.Name, Atharva.Email)

	Atharva.GetStatus()

	// won't change the email, as pass by value and we are not using pointers
	Atharva.NewMail()

	fmt.Printf("Name is %v and Email is %v\n", Atharva.Name, Atharva.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//  this is a method, the function which is associated with a struct
func (u User) GetStatus(){
	fmt.Println("Is user active? : ", u.Status)
}

// wont change the email, as pass by value and we are not using pointers
func (u User) NewMail(){
	u.Email = "test@gmai.com"
	fmt.Println("Email of this user is : ", u.Email)
}