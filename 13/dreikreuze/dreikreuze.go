package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type GameStatus int

const (
	running GameStatus = iota
	quit
	humanWins
	computerWins
)

const boardWidth = 23

type Player interface {
	makeMove(*uint) GameStatus
}

type ComputerPlayer struct{}

type HumanPlayer struct{}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	intro()
	startGame()
}

func intro() {
	fmt.Println("*** Tic Tac Toe ***")
	fmt.Println("Given a chain of 23 free fields. Each player takes turns placing an X on an empty field.")
	fmt.Println("If three or more X's are adjacent, the player wins.")
}

func startGame() {
	var board uint
	var currentPlayer, nextPlayer Player
	currentPlayer = ComputerPlayer{}
	nextPlayer = HumanPlayer{}

	gameStatus := running
	for gameStatus == running {
		printBoard(board)
		currentPlayer, nextPlayer = nextPlayer, currentPlayer
		gameStatus = currentPlayer.makeMove(&board)
		if gameStatus == quit {
			break
		}
		gameStatus = checkGameStatus(board, currentPlayer)
	}
	printBoard(board)
	printGameResult(gameStatus)
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

func (c ComputerPlayer) makeMove(board *uint) GameStatus {
	for {
		move := uint(rand.Intn(int(boardWidth)))
		if !isOccupied(move, *board) {
			occupy(move, board)
			fmt.Printf("Computer occupies field %d\n", move)
			return running
		}
	}
}

func (h HumanPlayer) makeMove(board *uint) GameStatus {
	for {
		fmt.Print("Enter your move (0-22) or 'q' to quit: ")
		var input string
		fmt.Scanln(&input)
		if input == "q" {
			return quit
		}
		move, err := strconv.ParseUint(input, 10, 32)
		if err == nil && move < boardWidth && !isOccupied(uint(move), *board) {
			occupy(uint(move), board)
			return running
		}
		fmt.Println("Invalid move. Try again.")
	}
}

func checkGameStatus(board uint, currentPlayer Player) GameStatus {
	crosses := 0
	for i := 0; i < boardWidth; i++ {
		if isOccupied(uint(i), board) {
			crosses++
			if crosses == 3 {
				if _, ok := currentPlayer.(HumanPlayer); ok {
					return humanWins
				}
				return computerWins
			}
		} else {
			crosses = 0
		}
	}
	return running
}

func printGameResult(gameStatus GameStatus) {
	switch gameStatus {
	case quit:
		fmt.Println("Player quits game")
	case humanWins:
		fmt.Println("Player wins!")
	case computerWins:
		fmt.Println("Computer wins!")
	}
}

func occupy(pos uint, board *uint) {
	bit := uint(1) << pos
	*board |= uint(bit)
}

func isOccupied(pos uint, board uint) bool {
	bit := uint(1) << pos
	return (bit & uint(board)) == bit
}
