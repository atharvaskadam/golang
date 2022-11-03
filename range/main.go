package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 6, 7}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)
	for i, num := range nums {
		if num == 5 {
			fmt.Println("Index:", i)
		}
	}
	fruit := map[string]string{"a": "Apple", "b": "Orange"}
	for k, v := range fruit {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k:=range fruit{
		fmt.Println("Key",k)
	}
	for i,c:=range "ASK"{
		fmt.Println(i,c)
	}
}
