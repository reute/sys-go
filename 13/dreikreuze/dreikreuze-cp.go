package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type GameStatus int

const (
	running GameStatus = iota
	quit
	playerWins
	computerWins
)

type Player interface {
	move(board *Board) GameStatus
	isHuman() bool
}

type ComputerPlayer struct {
	name string
}

type HumanPlayer struct {
	name string
}

const boardWidth = 23

type Board uint

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	startGame()
}

func startGame() {
	intro()
	var board Board
	hplayer := HumanPlayer{name: "RUDI"}
	cplayer := ComputerPlayer{name: "PETER"}
	var nextPlayer Player

	var start string
	fmt.Print("Do you want to start? (y/n): ")
	fmt.Scanf("%s", &start)
	if strings.ToLower(start) == "y" {
		nextPlayer = &cplayer
	} else {
		nextPlayer = &hplayer
	}

	gameStatus := running
	printBoard(board)
	for gameStatus == running {
		nextPlayer = getNextPlayer(nextPlayer, &hplayer, &cplayer)
		move := nextPlayer.move(&board)
		if move == quit {
			break
		}
		printBoard(board)
		gameStatus = checkGameStatus(board, nextPlayer)
	}
	printGameResult(gameStatus)
}
func getNextPlayer(currentPlayer, hplayer, cplayer Player) Player {
	if currentPlayer == hplayer {
		return cplayer
	} else {
		return hplayer
	}
}

func intro() {
	fmt.Println("*** Tic Tac Toe ***")
	fmt.Println("Given a chain of 23 free fields. Each player takes turns placing an X on an empty field.")
	fmt.Println("If three or more X's are adjacent, the player wins.")
}

func printGameResult(gameStatus GameStatus) {
	switch gameStatus {
	case quit:
		fmt.Println("Player quits game")
	case playerWins:
		fmt.Println("Player wins!")
	case computerWins:
		fmt.Println("Computer wins!")
	}
}

func (b *Board) occupy(pos uint) {
	bit := uint(1) << pos
	*b |= Board(bit)
}

func (b Board) isOccupied(pos uint) bool {
	bit := uint(1) << pos
	return (bit & uint(b)) == bit
}

func checkGameStatus(b Board, currentPlayer Player) GameStatus {
	crosses := 0
	for i := 0; i < boardWidth; i++ {
		if b.isOccupied(uint(i)) {
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

func (c *ComputerPlayer) move(board *Board) GameStatus {
	for {
		move := uint(rand.Intn(int(boardWidth)))
		if !board.isOccupied(move) {
			board.occupy(move)
			fmt.Printf("Computer occupies field %d\n", move)
			return running
		}
	}
}

func (h *HumanPlayer) move(board *Board) GameStatus {
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
		if board.isOccupied(move) {
			fmt.Println("Field is already occupied!")
			continue
		}
		board.occupy(move)
		fmt.Printf("Player occupies field %d\n", move)
		return running
	}
}

func (HumanPlayer) isHuman() bool {
	return true
}

func (ComputerPlayer) isHuman() bool {
	return false
}

func printBoard(board Board) {
	for i := 0; i < boardWidth; i++ {
		if board.isOccupied(uint(i)) {
			fmt.Print("\\/ ")
		} else {
			fmt.Print("   ")
		}
	}
	fmt.Println()
	for i := 0; i < boardWidth; i++ {
		if board.isOccupied(uint(i)) {
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
