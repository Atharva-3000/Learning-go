package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Go class")

	var fruitList = []string{"Apple", "Orange", "Banana"}

	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Guava")

	fmt.Println("Fruit List is: ", fruitList)

	fruitList = fruitList[1:3] // [Orange, Banana]

	fmt.Println("Fruit List after slicing: ", fruitList)

	highScore := make([]int, 4)  //4 is number of values and not capacity

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 777

	highScore = append(highScore, 555, 666, 777) // 4 more values added to highScore as after defafult allocation
	fmt.Println("High Score is: ", highScore)

	sort.Ints(highScore)
	fmt.Println("Sorted High Score is: ", highScore)

	fmt.Println(sort.IntsAreSorted(highScore))
}
