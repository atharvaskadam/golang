package main

import "fmt"

func main() {
	i := 2
	fmt.Print("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	a := "BMW"
	fmt.Print(a, " is ")
	switch a {
	case a:
		fmt.Println("Car")
	case a:
		fmt.Println("Fruit")
	}
}
