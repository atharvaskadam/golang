package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}
func plusPlus(a, b, c int) int {
	return a + b + c
}
func mul(x, y, z int) int {
	return x * y * z
}
func main() {
	res := plus(2, 5)
	fmt.Println("2+5=", res)

	res = plusPlus(5, 6, 7)
	fmt.Println("5+6+7=", res)
	ans := mul(4, 4, 4)
	fmt.Println("Multiplication Is", ans)
}
