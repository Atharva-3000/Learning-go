package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study in Go")

	presentTime := time.Now()
	fmt.Println("The current time is: ", presentTime)

	// weird format

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.July, 22, 23, 0, 0, 0, time.UTC).Format("02-01-2006 15:04:05 Monday")

	fmt.Println("The created date is: ", createdDate)
}
