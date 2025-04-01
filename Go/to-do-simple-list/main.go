package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tasks := make([]string, 0)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Essa é uma to-do list simples.")
	fmt.Println("Comandos disponíveis: add, list, exit")

	for {
		fmt.Println("\n> ")
		input, _ := reader.ReadString('\n')
		cmd := strings.TrimSpace(input)

		switch cmd {
		case "add":
			fmt.Print("Digite a tarefa que deseja adicionar: ")
			task, _ := reader.ReadString('\n')
			tasks = append(tasks, strings.TrimSpace(task))
			fmt.Println("Tarefa adicionada:", task)

		case "list":
			if len(tasks) == 0 {
				fmt.Println("Nenhuma tarefa encontrada.")
			} else {
				fmt.Println("Tarefas:")
				for i, task := range tasks {
					fmt.Printf("%d: %s\n", i+1, task)
				}
			}

		case "exit":
			fmt.Println("Até logo!")
			return

		default:
			fmt.Println("Comando inválido. Tente novamente.")
		}
	}
}
