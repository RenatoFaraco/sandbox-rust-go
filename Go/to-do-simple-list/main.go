package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Códigos ANSI para cores e formatação
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Bold    = "\033[1m"
)

func colorize(text string, colorCode string) string {
	return colorCode + text + Reset
}

func printHeader() {
	header := `
████████╗ ██████╗ ██████╗  ██████╗ 
╚══██╔══╝██╔═══██╗██╔══██╗██╔═══██╗
   ██║   ██║   ██║██║  ██║██║   ██║
   ██║   ██║   ██║██║  ██║██║   ██║
   ██║   ╚██████╔╝██████╔╝╚██████╔╝
   ╚═╝    ╚═════╝ ╚═════╝  ╚═════╝ 
`
	fmt.Println(colorize(header, Cyan))
	fmt.Println(colorize("Bem-vindo à sua To-Do List Terminal!", Yellow+Bold))
	fmt.Println(colorize("Digite 'help' para ver os comandos disponíveis", Blue))
	fmt.Println(strings.Repeat("─", 60))
}

func main() {
	tasks := make([]string, 0)
	reader := bufio.NewReader(os.Stdin)

	printHeader()

	for {
		fmt.Print(colorize("\n❯ ", Green+Bold))
		input, _ := reader.ReadString('\n')
		cmd := strings.TrimSpace(input)

		switch cmd {
		case "add":
			fmt.Print(colorize("✏  Digite a tarefa: ", Magenta))
			task, _ := reader.ReadString('\n')
			cleanTask := strings.TrimSpace(task)
			tasks = append(tasks, cleanTask)
			fmt.Println(colorize("✓ Tarefa adicionada: "+cleanTask, Green))

		case "remove":
			if len(tasks) == 0 {
				fmt.Println(colorize("⚠  Nenhuma tarefa para remover.", Red))
				continue
			}
			fmt.Print(colorize("✂  Número da tarefa a remover (1-"+strconv.Itoa(len(tasks))+"): ", Magenta))
			input, _ := reader.ReadString('\n')
			index, err := strconv.Atoi(strings.TrimSpace(input))

			if err != nil || index < 1 || index > len(tasks) {
				fmt.Println(colorize("⚠  Número inválido!", Red))
				continue
			}
			removedTask := tasks[index-1]
			tasks = append(tasks[:index-1], tasks[index:]...)
			fmt.Println(colorize("🗑  Tarefa removida: "+removedTask, Yellow))

		case "clear":
			tasks = make([]string, 0)
			fmt.Println(colorize("♻  Todas as tarefas foram removidas!", Yellow))

		case "list":
			fmt.Println(colorize("\n📋 Lista de Tarefas:", Bold))
			if len(tasks) == 0 {
				fmt.Println(colorize("~ Nenhuma tarefa encontrada ~", Blue))
			} else {
				for i, task := range tasks {
					fmt.Printf(colorize(" %d ➔ ", Cyan+Bold), i+1)
					fmt.Println(colorize(task, Reset))
				}
			}
			fmt.Println(strings.Repeat("─", 30))

		case "help":
			fmt.Println(colorize("\n📚 Comandos Disponíveis:", Bold))
			fmt.Println(colorize("  add    ", Green) + colorize("- Adiciona nova tarefa", Reset))
			fmt.Println(colorize("  remove ", Yellow) + colorize("- Remove uma tarefa", Reset))
			fmt.Println(colorize("  clear  ", Red) + colorize("- Limpa todas as tarefas", Reset))
			fmt.Println(colorize("  list   ", Blue) + colorize("- Lista todas as tarefas", Reset))
			fmt.Println(colorize("  help   ", Magenta) + colorize("- Mostra esta ajuda", Reset))
			fmt.Println(colorize("  exit   ", Cyan) + colorize("- Sai do programa", Reset))

		case "exit":
			fmt.Println(colorize("\n🖖 Até logo! Volte sempre :)", Cyan+Bold))
			os.Exit(0)

		default:
			fmt.Println(colorize("⚠  Comando inválido. Digite 'help' para ajuda.", Red))
		}
	}
}
