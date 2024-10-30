package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println(("Post request in Go"))
	PostRequest()

}

func PostRequest() {
	// POST request
	const myurl = "http://localhost:3000/post"

	// fake json payload

	requestBody := strings.NewReader(`{
	"title":"New Post 1","content":"The new post content 1",
	"author":"John Doe"
	}
	`)

	response, error := http.Post(myurl, "application/json", requestBody)
	handleErr(error)
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	// fmt.Println(string(content))
	responseString.Write(content)
	fmt.Println(responseString.String())
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
