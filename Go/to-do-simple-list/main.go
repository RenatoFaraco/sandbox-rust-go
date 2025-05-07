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

	fmt.Print("Essa é uma to-do list simples.")
	fmt.Print("Comandos disponíveis: add, list, exit")

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

		case "remove":
			fmt.Print("Digite o número da tarefa que deseja remover: ")
			var index int
			_, err := fmt.Scanf("%d", &index)
			if err != nil || index < 1 || index > len(tasks) {
				fmt.Println("Número inválido.")
				continue
			}
			tasks = append(tasks[:index-1], tasks[index:]...)
			fmt.Println("Tarefa removida:", tasks[index-1])

		case "clear":
			tasks = make([]string, 0)
			fmt.Println("Todas as tarefas foram removidas.")

		case "list":
			if len(tasks) == 0 {
				fmt.Println("Nenhuma tarefa encontrada.")
			} else {
				fmt.Println("Tarefas:")
				for i, task := range tasks {
					fmt.Printf("%d: %s\n", i+1, task)
				}
			}

		case "help":
			fmt.Println("Comandos disponíveis:")
			fmt.Println("add - Adiciona uma nova tarefa")
			fmt.Println("remove - Remove uma tarefa existente")
			fmt.Println("clear - Remove todas as tarefas")
			fmt.Println("list - Lista todas as tarefas")
			fmt.Println("help - Mostra os comandos disponíveis")
			fmt.Println("exit - Sai do programa")

		case "exit":
			fmt.Println("Até logo!")
			return

		default:
			fmt.Println("Comando inválido. Tente novamente.")
		}
	}
}
