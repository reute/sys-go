package main

import (
	"fmt"
	"math/rand"
)

const (
	player = iota
	comp
)

func main() {
	inputPlayer := func(number int, valid func(int, int) bool) int {
		for {
			fmt.Print("Your move: ")
			var move int
			if _, err := fmt.Scanf("%d", &move); err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			if valid(move, number) {
				return move
			}
		}
	}

	inputComputer := func(number int, valid func(int, int) bool) int {
		for {
			move := rand.Intn(number-1) + 1
			if valid(move, number) {
				fmt.Println("My move:", move)
				return move
			}
		}
	}

	isValidMove := func(move, number int) bool {
		if move == 0 || (number%move != 0 && move != number) {
			fmt.Printf("%d is not a proper divisor of %d\n", move, number)
			return false
		}
		return true
	}

	funcArray := []func(int, func(int, int) bool) int{inputPlayer, inputComputer}

	number := rand.Intn(100) + 1 // Ensure number is at least 1
	turn := player

	fmt.Println("*** Aliquot Game ***")
	for {
		fmt.Println("Number:", number)
		move := funcArray[turn](number, isValidMove)
		number -= move
		if gameover(number) {
			printResult(uint(turn))
			break
		}
		turn = (turn + 1) % 2
	}
}

func printResult(turn uint) {
	if turn == comp {
		fmt.Println("You lose !")
	} else if turn == player {
		fmt.Println("You win !")
	}
}

func gameover(number int) bool {
	if number == 1 {
		return true
	} else {
		return false
	}
}
