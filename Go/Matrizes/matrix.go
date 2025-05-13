package main

import (
	"fmt"
)

//notação : [n]T matriz de tipo T de tamanho n
//ex: var a [10]int matriz a de 10 inteiros

func main() {
	var a [3]string
	a[0] = "Brinca"
	a[1] = "com"
	a[2] = "matrizes"

	for i := 0; i < 3; i++ {
		fmt.Println(a[i])
	}

	primos := [6]int{1, 3, 5, 7, 11, 13}
	fmt.Println(primos)

	//Uma slice, é dinamicamente redimensionada,
	// uma visão flexível dos elementos de uma matriz.
	var s []int = primos[0:2]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	c := names[0:2]
	d := names[1:]
	fmt.Println(c, d)
	printSlice(c)

	d[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 12}
	fmt.Println(q)

	r := []bool{true, false, false, true}
	fmt.Println(r)

	s2 := []struct {
		x int
		y bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)
}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}
