package main

import (
	"fmt"
)

func main() {
	fmt.Println("...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Feito")
}
