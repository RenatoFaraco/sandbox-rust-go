package main

import "fmt"

func main() {
	q := []int{1, 2, 3, 4, 5}
	fmt.Println(q)

	r := []bool{true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
		s string
	}{
		{1, true, "one"},
		{2, false, "two"},
		{3, true, "three"},
		{4, false, "four"},
	}
	fmt.Println(s)

}
