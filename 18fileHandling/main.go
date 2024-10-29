package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("File Handling in Go")
	content := "This needs to go inside the file - Atharva"

	file, err := os.Create("./sample.txt")
	checkNilError(err)
	length, error := io.WriteString(file, content)
	checkNilError(error)
	fmt.Println("Length of the file is: ", length)
	defer file.Close()
	readFile("./sample.txt")
}

func readFile(fileName string)  {
	dataBytes, error := ioutil.ReadFile(fileName)
	checkNilError(error)
	fmt.Println(string(dataBytes))
}


func checkNilError(err error){
	if err != nil {
		panic(err)
	}
}