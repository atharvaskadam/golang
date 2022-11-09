package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum", sum)

	nums1 := []int{2, 3, 4}
	num1 := 0
	for _, num := range nums1 {
		sum += num
	}
	fmt.Println("Sum", num1)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}
	fruit := map[string]string{"a": "Apple", "b": "Orange"}
	for k, i := range fruit {
		fmt.Printf("%s => %s \n", k, i)
	}
	for k := range fruit {
		fmt.Println("Key", k)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
