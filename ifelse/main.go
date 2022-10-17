package main

import (
	"fmt"
)

func main() {
	if 8%2 == 0 {
		fmt.Println("Output Is Even")
	} else {
		fmt.Println("Output Is Odd")
	}
	if 10%5 == 2 {
		fmt.Println("Remainder  is 2")
	} else {
		fmt.Println("Remainder is Something Else")
	}
}
