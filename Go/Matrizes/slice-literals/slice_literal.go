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
		{1, true, "1"},
		{2, false, "2"},
		{3, true, "3"},
		{4, false, "4"},
	}
	fmt.Println(s)

}
