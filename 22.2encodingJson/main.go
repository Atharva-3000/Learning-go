package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"`  // this will not be encoded not be reflected in json
	Tags     []string `json:"tags,omitempty"` // if tags is empty, it will not be reflected in json
}

func main() {
	fmt.Println(("Json creation ini Go"))
	encodeJson()
}

func encodeJson() {
	myCourses := []course{
		{"ReactJS Bootcamp", 100, "Portal", "abc@123", []string{"React", "Javascript", "Node", "Web"}},

		{"MERN Bootcamp", 900, "Portal", "abc@def", []string{"React", "Javascript", "Node", "Web", "Databse", "Full Stack"}},
	}

	// package this data as json data
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")
	errHandler(err)
	fmt.Printf("%s\n", finalJson)
}

func errHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
