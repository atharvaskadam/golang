package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[1])

	fmt.Println("len:", len(s))

	fmt.Println("----------------------------")
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("appended:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	fmt.Println("----------------------------")
	t := make([]string, 4)
	fmt.Println("car:", t)

	t[0] = "Maruti"
	t[1] = "BMW"
	t[2] = "Audi"
	t[3] = "Bugati"
	fmt.Println("Set", t)
	fmt.Println("Get", t[3])
	fmt.Println("len", len(t))
	t = append(t, "Nissan")
	t = append(t, "Skoda", "Ferrari")
	fmt.Println("Appended slice", t)

	u := make([]string, len(t))
	copy(u, t)
	fmt.Println("cpy", u)

	l := s[2:5]
	fmt.Println("Sli", l)

	f := s[0:4]
	fmt.Println("Slice", f)
}
