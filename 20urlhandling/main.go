package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://atharva-builds:3000/vercel?name=atharva&age=21"

func main() {
	fmt.Println("URL Handling in Go")
	fmt.Println("URL is: ", myUrl)

	result, err := url.Parse(myUrl)
	handleError(err)
	// fmt.Println(result. Scheme);
	// fmt.Println(result.Host);
	// fmt.Println(result.Path);
	// fmt.Println(result.RawQuery);
	// fmt.Println(result.Port());

	// Query Params returns a map of query parameters
	qparams := result.Query();
	fmt.Printf("Query Params: %v\n", qparams);
	// fmt.Println(qparams["name"]);
	// fmt.Println(qparams["age"]);

	for _, value := range qparams {
		fmt.Println(value);
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "atharva-builds:3000",
		Path: "/vercel",
		RawQuery: "name=atharva&age=21",
	}

	anotherUrl := partsOfUrl.String();
	fmt.Println("Another URL: ", anotherUrl);
}

func handleError (err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}