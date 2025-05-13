package main

import (
	"fmt"
)

type Exemplo struct {
	X int
	Y int
}

var (
	e1 = Exemplo{1, 2}
	e2 = Exemplo{X: 1}
	e3 = Exemplo{}
	p  = &Exemplo{1, 2}
)

func main() {
	fmt.Println(e1, *p, e2, e3)
}
