package main

import (
	"fmt"
	"math/rand"
)

const (
	END = iota
	ONGOING
	BROKE
)

const NUM_DICES = 3

type t_dices [NUM_DICES]int

type Player struct {
	balance int
	name    string
}

type Game struct {
	player Player
}

func calcRoundResult(bet, guess int, dices t_dices) (win int) {
	for _, dice := range dices {
		fmt.Printf("%d ", dice)
		if dice == guess {
			win += bet
		}
	}
	if win == 0 {
		return -bet
	}
	return
}

func rollDices() (dices t_dices) {
	for i := range dices {
		dices[i] = rand.Intn(6) + 1
	}
	return
}

func inputGuess() (guess int) {
	for {
		fmt.Print("Your guess (1-6): ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil || guess < 1 || guess > 6 {
			fmt.Println("Invalid input. Please enter a number between 1 and 6.")
			continue
		}
		return
	}
}

func inputBet(balance int) (bet int) {
	for {
		fmt.Print("Your bet: ")
		_, err := fmt.Scanf("%d", &bet)
		if err != nil || bet < 0 || bet > balance {
			fmt.Println("Invalid bet. Please enter a positive number within your balance.")
			continue
		}
		return
	}
}

func newRound(player *Player) int {
	fmt.Printf("You have %d units\n", player.balance)
	bet := inputBet(player.balance)
	if bet == 0 {
		return END
	}
	guess := inputGuess()
	dices := rollDices()
	fmt.Print("The dice rolled: ")
	result := calcRoundResult(bet, guess, dices)
	if result <= 0 {
		fmt.Printf("Bad luck, no matches! You lose %d units\n", result)
	} else {
		fmt.Printf("Congratulations, you win %d units!\n", result)
	}
	player.balance += result
	if player.balance <= 0 {
		return BROKE
	}
	return ONGOING
}

func startGame(game *Game) {
	fmt.Println("**** Chuck-a-luck ****")
	for {
		status := newRound(&game.player)
		if status == END {
			fmt.Println("Game over. Thanks for playing!")
			break
		} else if status == BROKE {
			fmt.Println("You're broke. Game over!")
			break
		}
	}
}

func main() {
	player := Player{balance: 100, name: "Player"}
	game := Game{player: player}
	startGame(&game)
}
