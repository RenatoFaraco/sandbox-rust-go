package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	fmt.Println("JOGO DA VELHA - Terminal")
	fmt.Println("Instruções: Digite as coordenadas no formato LINHA,COLUNA (ex: 1,2)")

	for {
		printBoard(board)
		fmt.Printf("\nJogador %s, digite sua jogada: ", player)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		coords := strings.Split(input, ",")

		if len(coords) != 2 {
			fmt.Println("\nFormato inválido! Use LINHA,COLUNA (ex: 1,2)")
			continue
		}

		row, err1 := strconv.Atoi(coords[0])
		col, err2 := strconv.Atoi(coords[1])

		if err1 != nil || err2 != nil {
			fmt.Println("\nCoordenadas devem ser números!")
			continue
		}

		row--
		col--

		if row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Println("\nPosição inválida! Use valores entre 1 e 3")
			continue
		}

		if board[row][col] != "_" {
			fmt.Println("\nPosição já ocupada!")
			continue
		}

		board[row][col] = player
		moves++

		if checkWinner(board, player) {
			printBoard(board)
			fmt.Printf("\nJogador %s venceu!\n", player)
			break
		}

		if moves == 9 {
			printBoard(board)
			fmt.Println("\nEmpate!")
			break
		}

		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
	}
}

func printBoard(board [][]string) {
	fmt.Println("\n  1 2 3")
	for i, row := range board {
		fmt.Printf("%d ", i+1)
		fmt.Println(strings.Join(row, " "))
	}
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
