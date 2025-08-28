package main

import "fmt"

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	} else if n == 1 {
		return []int{0}
	}

	seq := make([]int, n)
	seq[0] = 0
	seq[1] = 1

	for i := 2; i < n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}

	return seq
}

func main() {
	n := 100
	fibSeq := fibonacci(n)
	fmt.Print("Os números de Fibonacci --> \n")

	for i, num := range fibSeq {
		fmt.Printf("Fib(%d) = %d \n", i, num)
	}
}
