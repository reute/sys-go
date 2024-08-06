package main

import (
	"fmt"
	"math/rand"
)

const (
	boardWidth = iota + 23
	running
	quit
	playerWins
	computerWins
)

type Player interface {
	move(board *uint) (move uint)
	printMove(move uint)
}

type ComputerPlayer struct {
	name string
}

type HumanPlayer struct {
	name string
}

func main() {
	startGame()
}

func startGame() {
	fmt.Println("*** Tic Tac Toe ***")
	fmt.Println("Given a chain of 23 free fields. Each player takes turns placing an X on an empty field.")
	fmt.Println("If three or more X's are adjacent, the player wins.")

	hplayer := HumanPlayer{name: "RUDI"}
	cplayer := ComputerPlayer{name: "PETER"}
	var currentPlayer, nextPlayer Player

	var board uint
	var start string
	fmt.Print("Do you want to start? (y/n): ")
	fmt.Scanf("%s", &start)

	if start == "y" {
		currentPlayer = hplayer
		nextPlayer = cplayer
	} else {
		currentPlayer = cplayer
		nextPlayer = cplayer
	}

	gameStatus := running
	for gameStatus == running {
		printBoard(board)
		gameStatus = nextMove(currentPlayer, &board)
		currentPlayer, nextPlayer = nextPlayer, currentPlayer
	}
	printBoard(board)

	switch gameStatus {
	case quit:
		fmt.Println("Player quits")
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
	var crosses uint
	for i := uint(0); i < boardWidth; i++ {
		if isOccupied(i, board) {
			crosses++
		} else {
			crosses = 0
		}
		if crosses == 3 {
			_, ok := currentPlayer.(HumanPlayer)
			if ok {
				return playerWins
			} else {
				return computerWins
			}
		}
	}
	return running
}

func (ComputerPlayer) move(board *uint) (move uint) {
	for {
		move = uint(rand.Intn(int(boardWidth)))
		if !isOccupied(move, *board) {
			occupyField(move, board)
			return
		}
	}
}

func (HumanPlayer) move(board *uint) (move uint) {
	for {
		fmt.Print("Your move: ")
		fmt.Scanf("%d", &move)
		if move == 99 {
			return quit
		}
		if move < 0 || move >= boardWidth {
			fmt.Printf("Please enter a number between 0 and %d\n", boardWidth-1)
			continue
		}
		if isOccupied(move, *board) {
			fmt.Println("Field is already occupied!")
			continue
		}
		occupyField(move, board)
		return
	}
}

func (HumanPlayer) printMove(move uint) {
	fmt.Printf("Player occupies field %d\n", move)
}

func (ComputerPlayer) printMove(move uint) {
	fmt.Printf("Computer occupies field %d\n", move)
}

func nextMove(player Player, board *uint) int {
	move := player.move(board)
	if move == quit {
		return quit
	}
	player.printMove(move)
	return checkGameStatus(*board, player)
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
