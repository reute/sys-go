package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	boardWidth = iota + 23
	running
	quit
	playerWins
	computerWins
)

type Player interface {
	move() uint
	isHuman() bool
}

type ComputerPlayer struct {
	name  string
	board *uint
}

type HumanPlayer struct {
	name  string
	board *uint
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	startGame()
}

func startGame() {
	intro()
	var board uint
	hplayer := HumanPlayer{name: "RUDI", board: &board}
	cplayer := ComputerPlayer{name: "PETER", board: &board}
	var currentPlayer, nextPlayer Player

	var start string
	fmt.Print("Do you want to start? (y/n): ")
	fmt.Scanf("%s", &start)
	if strings.ToLower(start) == "y" {
		currentPlayer = &cplayer
		nextPlayer = &hplayer
	} else {
		currentPlayer = &hplayer
		nextPlayer = &cplayer
	}

	gameStatus := running
	printBoard(board)
	for gameStatus == running {
		currentPlayer, nextPlayer = nextPlayer, currentPlayer
		move := currentPlayer.move()
		if move == quit {
			break
		}
		printBoard(board)
		gameStatus = checkGameStatus(board, currentPlayer)
	}
	printGameResult(gameStatus)
}

func intro() {
	fmt.Println("*** Tic Tac Toe ***")
	fmt.Println("Given a chain of 23 free fields. Each player takes turns placing an X on an empty field.")
	fmt.Println("If three or more X's are adjacent, the player wins.")
}

func printGameResult(gameStatus int) {
	switch gameStatus {
	case quit:
		fmt.Println("Player quits game")
	case playerWins:
		fmt.Println("Player wins!")
	case computerWins:
		fmt.Println("Computer wins!")
	}
}

func occupyField(pos uint, board *uint) {
	bit := uint(1) << pos
	*board |= bit
}

func isOccupied(pos uint, board uint) bool {
	bit := uint(1) << pos
	return (bit & board) == bit
}

func checkGameStatus(board uint, currentPlayer Player) int {
	crosses := 0
	for i := 0; i < boardWidth; i++ {
		if isOccupied(uint(i), board) {
			crosses++
			if crosses == 3 {
				if currentPlayer.isHuman() {
					return playerWins
				}
				return computerWins
			}
		} else {
			crosses = 0
		}
	}
	return running
}

func (c *ComputerPlayer) move() uint {
	for {
		move := uint(rand.Intn(int(boardWidth)))
		if !isOccupied(move, *c.board) {
			occupyField(move, c.board)
			fmt.Printf("Computer occupies field %d\n", move)
			return move
		}
	}
}

func (h *HumanPlayer) move() uint {
	var input_pl int
	for {
		fmt.Print("Your move: ")
		fmt.Scanf("%d", &input_pl)
		if input_pl == 99 {
			return quit
		}
		if input_pl < 0 || input_pl >= boardWidth {
			fmt.Printf("Please enter a number between 0 and %d\n", boardWidth-1)
			continue
		}
		move := uint(input_pl)
		if isOccupied(move, *h.board) {
			fmt.Println("Field is already occupied!")
			continue
		}
		occupyField(move, h.board)
		fmt.Printf("Player occupies field %d\n", move)
		return move
	}
}

func (HumanPlayer) isHuman() bool {
	return true
}

func (ComputerPlayer) isHuman() bool {
	return false
}

func printBoard(board uint) {
	for i := 0; i < boardWidth; i++ {
		if isOccupied(uint(i), board) {
			fmt.Print("\\/ ")
		} else {
			fmt.Print("   ")
		}
	}
	fmt.Println()
	for i := 0; i < boardWidth; i++ {
		if isOccupied(uint(i), board) {
			fmt.Print("/\\ ")
		} else {
			fmt.Print("   ")
		}
	}
	fmt.Println()
	for i := 0; i < boardWidth; i++ {
		fmt.Printf("%02d ", i)
	}
	fmt.Println()
}
