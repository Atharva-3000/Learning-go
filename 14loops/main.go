package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	days := []string{
		"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday",
	}

	fmt.Println(days)
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// sorta kinda for each loop
	// for index, day := range days {
	// 	fmt.Printf("Day %v of the week is %v\n", index+1, day)
	// }

	rogueValue := 1

	for rogueValue < 10 {
		if(rogueValue==2){
			goto lco
		}
		if rogueValue == 5 {
			// break
			rogueValue++
			continue
		}
		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping to LCO")
}
