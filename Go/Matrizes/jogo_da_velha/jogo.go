package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Game struct {
	board     [][]string
	player    string
	moves     int
	score     map[string]int
	gameOver  bool
	lastMove  [2]int
	animating bool
}

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	BgWhite     = "\033[47m"
	BgBlue      = "\033[44m"
	BgRed       = "\033[41m"
)

func main() {
	game := Game{
		board: [][]string{
			{"_", "_", "_"},
			{"_", "_", "_"},
			{"_", "_", "_"},
		},
		player: "X",
		score:  map[string]int{"X": 0, "O": 0, "Empates.": 0},
	}

	clearScreen()
	showTitle()
	showInstructions()

	for {
		game.renderBoard()

		if game.gameOver {
			game.showGameResult()
			if askToPlayAgain() {
				game.reset()
				clearScreen()
				showTitle()
				continue
			}
			break
		}

		game.processMove()
	}

	fmt.Println("\nObrigado por jogar!")
}

func (g *Game) reset() {
	g.board = [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	g.player = "X"
	g.moves = 0
	g.gameOver = false
	g.lastMove = [2]int{-1, -1}
}

func (g *Game) renderBoard() {
	fmt.Print("\033[H\033[2J") // Clear screen without full refresh

	// Header with score
	fmt.Printf("%s=== JOGO DA VELHA ===%s", ColorCyan, ColorReset)
	fmt.Printf("\t%sX: %d  |  O: %d  |  Empates: %d%s\n",
		ColorYellow, g.score["X"], g.score["O"], g.score["Empates"], ColorReset)

	// Column numbers
	fmt.Printf("\n   %s1%s   %s2%s   %s3%s\n",
		ColorCyan, ColorReset, ColorCyan, ColorReset, ColorCyan, ColorReset)
	fmt.Println(" " + ColorYellow + "-------------" + ColorReset)

	for i, row := range g.board {
		// Row number
		fmt.Printf("%s%d%s", ColorCyan, i+1, ColorReset)

		for j, cell := range row {
			fmt.Print(ColorYellow + "|" + ColorReset + " ")

			// Highlight last move
			if g.lastMove[0] == i && g.lastMove[1] == j && !g.animating {
				fmt.Print(BgWhite + ColorBlue)
				if cell == "O" {
					fmt.Print(ColorRed)
				}
			}

			fmt.Printf(" %s ", g.getCellSymbol(cell))
			fmt.Print(ColorReset + " ")
		}

		fmt.Println(ColorYellow + "|" + ColorReset)
		if i < 2 {
			fmt.Println(" " + ColorYellow + "-------------" + ColorReset)
		}
	}

	fmt.Println(" " + ColorYellow + "-------------" + ColorReset)

	if !g.gameOver {
		fmt.Printf("\n%sJogador %s%s, digite sua jogada: ",
			ColorGreen, g.getPlayerSymbol(), ColorReset)
	}
}

func (g *Game) processMove() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		g.showError("Entrada vazia! Digite coordenadas.")
		return
	}

	if strings.Contains(input, " ") {
		input = strings.ReplaceAll(input, " ", ",")
	}

	coords := strings.Split(input, ",")
	if len(coords) < 2 {
		g.showError("Formato inválido! Use: LINHA,COLUNA (ex: 1,2)")
		return
	}

	row, err1 := strconv.Atoi(strings.TrimSpace(coords[0]))
	col, err2 := strconv.Atoi(strings.TrimSpace(coords[1]))

	if err1 != nil || err2 != nil {
		g.showError("Coordenadas devem ser números!")
		return
	}

	row--
	col--

	if row < 0 || row > 2 || col < 0 || col > 2 {
		g.showError("Posição inválida! Use valores entre 1 e 3")
		return
	}

	if g.board[row][col] != "_" {
		g.showError("Posição já ocupada!")
		return
	}

	// Animation effect
	g.animateMove(row, col)

	g.board[row][col] = g.player
	g.lastMove = [2]int{row, col}
	g.moves++

	if g.checkWinner() {
		g.score[g.player]++
		g.gameOver = true
		return
	}

	if g.moves == 9 {
		g.score["Empates"]++
		g.gameOver = true
		return
	}

	// Switch player
	if g.player == "X" {
		g.player = "O"
	} else {
		g.player = "X"
	}
}

func (g *Game) animateMove(row, col int) {
	g.animating = true
	original := g.board[row][col]

	for i := 0; i < 3; i++ {
		g.board[row][col] = "•"
		g.renderBoard()
		time.Sleep(100 * time.Millisecond)

		g.board[row][col] = "_"
		g.renderBoard()
		time.Sleep(100 * time.Millisecond)
	}

	g.board[row][col] = original
	g.animating = false
}

func (g *Game) showGameResult() {
	time.Sleep(500 * time.Millisecond)
	g.renderBoard()

	if g.checkWinner() {
		fmt.Printf("\n%s  === JOGADOR %s VENCEU! ===  %s\n",
			BgWhite+ColorBlue, g.getPlayerSymbol(), ColorReset)
	} else {
		fmt.Printf("\n%s  === EMPATE! ===  %s\n",
			BgWhite+ColorPurple, ColorReset)
	}

	fmt.Println()
}

func (g *Game) checkWinner() bool {
	player := g.player

	// Verifica linhas e colunas
	for i := 0; i < 3; i++ {
		if g.board[i][0] == player && g.board[i][1] == player && g.board[i][2] == player {
			return true
		}
		if g.board[0][i] == player && g.board[1][i] == player && g.board[2][i] == player {
			return true
		}
	}

	// Verifica diagonais
	if g.board[0][0] == player && g.board[1][1] == player && g.board[2][2] == player {
		return true
	}
	if g.board[0][2] == player && g.board[1][1] == player && g.board[2][0] == player {
		return true
	}

	return false
}

func (g *Game) getCellSymbol(cell string) string {
	switch cell {
	case "X":
		return ColorBlue + "X" + ColorReset
	case "O":
		return ColorRed + "O" + ColorReset
	case "•":
		return ColorYellow + "•" + ColorReset
	default:
		return " "
	}
}

func (g *Game) getPlayerSymbol() string {
	if g.player == "X" {
		return ColorBlue + "X" + ColorReset
	}
	return ColorRed + "O" + ColorReset
}

func (g *Game) showError(message string) {
	fmt.Printf("\n%sERRO: %s%s\n", ColorRed, message, ColorReset)
	fmt.Print("Pressione ENTER para continuar...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func askToPlayAgain() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nDeseja jogar novamente? (s/n): ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	return input == "s" || input == "sim"
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

func showTitle() {
	title := `
______         _           _   _             
| ___ \       | |         | | | |            
| |_/ /___  __| | ___  ___| |_| |_ ___ _ __  
|    // _ \/ _' |/ _ \/ __| __| __/ _ \ '__| 
| |\ \  __/ (_| | (_) \__ \ |_| ||  __/ |    
\_| \_\___|\__,_|\___/|___/\__|\__\___|_|    
                                             
			 _                              
			| |                             
			| | ___   __ _ _   _  ___ _ __  
			| |/ _ \ / _' | | | |/ _ \ '__| 
			| | (_) | (_| | |_| |  __/ |    
			|_|\___/ \__, |\__,_|\___|_|    
					|___/                   
`
	fmt.Println(ColorCyan + title + ColorReset)
}

func showInstructions() {
	fmt.Println(ColorYellow + "Instruções:" + ColorReset)
	fmt.Println(" - Digite coordenadas no formato LINHA,COLUNA (ex: 1,2)")
	fmt.Println(" - Use números de 1 a 3 para linha e coluna")
	fmt.Println(" - Pressione CTRL+C para sair a qualquer momento")
	fmt.Println(ColorYellow + "\nPlacar mostra vitórias de cada jogador e empates" + ColorReset)
	fmt.Println("\n" + ColorYellow + "Pressione ENTER para começar..." + ColorReset)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
