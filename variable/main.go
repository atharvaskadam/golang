package main

import "fmt"

func main() {
	//Normal Variable Declaration
	var a int = 12
	var b int = 13
	fmt.Println("Sum Of Variables Is", a+b)

	//Shorthand Declaration Can Be Done
	d := 34
	c := 45
	fmt.Println("Sum Of Shorthand Variables Is", d+c)

	//Boolean
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
