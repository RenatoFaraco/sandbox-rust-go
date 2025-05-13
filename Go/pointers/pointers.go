package main

import (
	"fmt"
)

func main() {
	// Temos as variáveis
	i, j := 10, 2000
	p := &i                                                   //ponteiro que aponta para i
	fmt.Println("Lê o valor de i através do pointeiro: ", *p) // Lê o valor de i através do pointeiro
	*p = 21                                                   // atribui um novo valor para i através do ponteiro
	fmt.Println("Novo valor de i: ", i)
	fmt.Println("O endereço que o ponteiro aponta é: ", p)

	p = &j       // o ponteiro passa a apontar para j
	*p = *p / 10 // Divide o valor de j por meio do ponteiro
	fmt.Println("Lê o valor de j através do pointeiro: ", *p)
	fmt.Println("O endereço que o ponteiro aponta: ", p)
}
