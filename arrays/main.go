package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("Emp", a)

	var b [6]string
	fmt.Println("Names", b)

	a[2] = 100
	fmt.Println("Emp", a)
	fmt.Println("Get", a[2])

	fmt.Println("Len", len(a))

	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println(c)
}
