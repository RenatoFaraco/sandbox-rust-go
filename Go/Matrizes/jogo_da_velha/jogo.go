package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func main() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	player := "X"
	moves := 0
	reader := bufio.NewReader(os.Stdin)

	clearScreen()
	fmt.Println(ColorCyan + "=== JOGO DA VELHA ===" + ColorReset)
	fmt.Println(ColorYellow + "Instruções: Digite as coordenadas no formato LINHA,COLUNA (ex: 1,2)")
	fmt.Println("Pressione CTRL+C para sair a qualquer momento" + ColorReset)
	fmt.Println()

	for {
		printBoard(board)
		fmt.Printf("\n%sJogador %s%s, digite sua jogada: ", ColorGreen, getPlayerSymbol(player), ColorReset)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Permitir diferentes formatos de entrada
		if strings.Contains(input, " ") {
			input = strings.ReplaceAll(input, " ", ",")
		}

		coords := strings.Split(input, ",")
		if len(coords) < 2 {
			showError("Formato inválido! Use: LINHA,COLUNA (ex: 1,2)")
			continue
		}

		row, err1 := strconv.Atoi(strings.TrimSpace(coords[0]))
		col, err2 := strconv.Atoi(strings.TrimSpace(coords[1]))

		if err1 != nil || err2 != nil {
			showError("Coordenadas devem ser números!")
			continue
		}

		row--
		col--

		if row < 0 || row > 2 || col < 0 || col > 2 {
			showError("Posição inválida! Use valores entre 1 e 3")
			continue
		}

		if board[row][col] != "_" {
			showError("Posição já ocupada!")
			continue
		}

		board[row][col] = player
		moves++

		if checkWinner(board, player) {
			clearScreen()
			printBoard(board)
			fmt.Printf("\n%s=== JOGADOR %s VENCEU! ===%s\n\n", ColorYellow, getPlayerSymbol(player), ColorReset)
			break
		}

		if moves == 9 {
			clearScreen()
			printBoard(board)
			fmt.Printf("\n%s=== EMPATE! ===%s\n\n", ColorPurple, ColorReset)
			break
		}

		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}

		clearScreen()
	}
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printBoard(board [][]string) {
	fmt.Println(ColorCyan + "    1   2   3" + ColorReset)
	fmt.Println("  " + ColorYellow + "-------------" + ColorReset)

	for i, row := range board {
		fmt.Printf(ColorCyan+"%d "+ColorReset, i+1)
		fmt.Print(ColorYellow + "| " + ColorReset)

		for j, cell := range row {
			fmt.Printf(" %s ", getCellSymbol(cell))
			if j < 2 {
				fmt.Print(ColorYellow + "| " + ColorReset)
			}
		}

		fmt.Println(ColorYellow + "|" + ColorReset)
		if i < 2 {
			fmt.Println("  " + ColorYellow + "-------------" + ColorReset)
		}
	}

	fmt.Println("  " + ColorYellow + "-------------" + ColorReset)
}

func getCellSymbol(cell string) string {
	switch cell {
	case "X":
		return ColorBlue + "X" + ColorReset
	case "O":
		return ColorRed + "O" + ColorReset
	default:
		return " "
	}
}

func getPlayerSymbol(player string) string {
	if player == "X" {
		return ColorBlue + "X" + ColorReset
	}
	return ColorRed + "O" + ColorReset
}

func showError(message string) {
	fmt.Printf("\n%sERRO: %s%s\n\n", ColorRed, message, ColorReset)
	fmt.Print("Pressione ENTER para continuar...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	clearScreen()
}

func checkWinner(board [][]string, player string) bool {
	// Verifica linhas e colunas
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}

	// Verifica diagonais
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}
