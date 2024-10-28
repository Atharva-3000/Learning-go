package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)
	languages["GO"] = "Golang"
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of languages: ", languages)

	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of languages after deleting: ", languages)


	// loops in maps
	for key, value := range languages {
		fmt.Printf("For Key: %v, Value: %v\n", key, value)
	}
}
