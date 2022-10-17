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
	if 10%5 == 0 {
		fmt.Println("Remainder  is 0")
	} else {
		fmt.Printf("Remainder is %d", 10%5)
	}
}
