package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go class")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Banana"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Onion"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Length of Veg list is: ", len(vegList))
}