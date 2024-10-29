package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Web Requests in Go")

	res, error := http.Get("https://jsonplaceholder.typicode.com/todos")
	checkNilError(error)
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	checkNilError(err)
	fmt.Println(string(data))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
