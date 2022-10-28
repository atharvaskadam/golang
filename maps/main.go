package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 44
	m["k2"] = 45
	m["k3"] = 67
	fmt.Println("Map", m)

	fmt.Println("================================")

	n := make(map[string]int)
	n["key1"] = 45
	n["key2"] = 40
	n["key3"] = 48
	n["key4"] = 60
	fmt.Println("Map", n)

	fmt.Println("================================")

	name := make(map[string]int)
	name["Akash"] = 56
	name["Atharva"] = 57
	name["Vivek"] = 67
	name["Sushant"] = 98
	fmt.Println("Names Are", name)

	fmt.Println("================================")

	v1 := m["k1"]  //Here The Value For key Is Fetched With "name[key]" i.e. k1
	v2 := name["Atharva"] 
	fmt.Println("V1", v1)
	fmt.Println("V2", v2)
}
