package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func writeFile() {
	file, err := os.Create("Student.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString("List Of Students")
	if err != nil {
		panic(err)
	}
}
func readFile() {
	data, err := ioutil.ReadFile("Student.txt")
	data, err = os.ReadFile("Student.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
func main() {
	writeFile()
	readFile()
}