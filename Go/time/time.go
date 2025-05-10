package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Olá....")
	fmt.Println("Agora é: ", time.Now().Format("15:04:05 02/01/2006"))
}
