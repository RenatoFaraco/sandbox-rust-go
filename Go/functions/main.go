package main

import (
	"fmt"
)

var c, python, java bool

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Print: ", add(1, 2))

	a, b := swap("Teste", "Função")
	fmt.Println("\n Print: ", a, b)

	fmt.Println(split(17))

	var i, J int = 1, 2
	c, python, java := true, false, "no!"
	fmt.Println("\n Print: ", i, J, c, python, java)

	var x int
	var y float64
	var z bool
	var w string
	fmt.Printf("%v %v %v %q\n", x, y, z, w)

	v := 42 + 52i
	fmt.Printf("v is of type %T\n", v)
	fmt.Println(v)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(needFloat(Small))

}
