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

const (
	NUM_DICES       = 3
	INITIAL_BALANCE = 100
	MIN_DICE_NUMBER = 1
	MAX_DICE_NUMBER = 6
)

type t_dicesValues [NUM_DICES]int

type Player struct {
	balance int
	name    string
}

func main() {
	name := inputPlayerName()
	player := Player{balance: INITIAL_BALANCE, name: name}
	startGame(&player)
}

func inputPlayerName() string {
	var name string
	for {
		fmt.Print("Enter your name: ")
		_, err := fmt.Scanf("%s", &name)
		if err != nil || name == "" {
			fmt.Println("Invalid name. Please try again.")
			continue
		}
		return name
	}
}

func startGame(player *Player) {
	fmt.Println("**** Chuck-a-luck ****")
	displayRules()
	roundsPlayed := 0
	initialBalance := player.balance
	status := ONGOING
	for status == ONGOING {
		status = playRound(player)
		roundsPlayed++
	}
	if status == BROKE {
		fmt.Println("You're broke. Game over!")
	} else {
		displayGameSummary(player, roundsPlayed, initialBalance)
	}
}

func displayRules() {
	rules := `
**** Chuck-a-luck Rules ****
1. You start with 100 units.
2. Each round, you bet on a number (1-6).
3. Three dice are rolled.
4. You win your bet for each die matching your guess.
5. If no dice match, you lose your bet.
6. The game ends when you're broke or choose to quit.
****************************************************`
	fmt.Println(rules)
}

func playRound(player *Player) int {
	fmt.Printf("You have %d units\n", player.balance)
	bet := inputBet(player.balance)
	if bet == 0 {
		return END
	}
	guess := inputGuess()
	dices := rollDices()
	result := calcResult(bet, guess, dices)
	player.balance += result
	if player.balance == 0 {
		return BROKE
	}
	return ONGOING
}

func inputBet(balance int) (bet int) {
	for {
		fmt.Print("Your bet: ")
		_, err := fmt.Scanf("%d", &bet)
		if err != nil || bet < 0 || bet > balance {
			fmt.Printf("Invalid bet. Please enter a positive number not bigger than %d\n", balance)
			continue
		}
		return
	}
}

func inputGuess() (guess int) {
	for {
		fmt.Printf("Your guess (%d-%d): ", MIN_DICE_NUMBER, MAX_DICE_NUMBER)
		_, err := fmt.Scanf("%d", &guess)
		if err != nil || guess < MIN_DICE_NUMBER || guess > MAX_DICE_NUMBER {
			fmt.Printf("Invalid input. Please enter a number between %d and %d\n", MIN_DICE_NUMBER, MAX_DICE_NUMBER)
			continue
		}
		return
	}
}

func rollDices() (dices t_dicesValues) {
	for i := range dices {
		dices[i] = rand.Intn(MAX_DICE_NUMBER) + MIN_DICE_NUMBER
	}
	return
}

func calcResult(bet, guess int, dices t_dicesValues) int {
	fmt.Print("The dice rolled: ")
	matches := 0
	for _, dice := range dices {
		fmt.Printf("%d ", dice)
		if dice == guess {
			matches++
		}
	}
	fmt.Println()
	if matches == 0 {
		return -bet
	}
	result := bet * matches
	if result <= 0 {
		fmt.Printf("Bad luck, no matches! You lose %d units\n", -result)
	} else {
		fmt.Printf("Congratulations, you win %d units!\n", result)
	}
	return result
}

func displayGameSummary(player *Player, roundsPlayed int, initialBalance int) {
	fmt.Printf("\n**** Game Summary for %s ****\n", player.name)
	fmt.Printf("Rounds played: %d\n", roundsPlayed)
	fmt.Printf("Initial balance: %d\n", initialBalance)
	fmt.Printf("Final balance: %d\n", player.balance)
	profit := player.balance - initialBalance
	if profit > 0 {
		fmt.Printf("Total profit: %d\n", profit)
	} else {
		fmt.Printf("Total loss: %d\n", -profit)
	}
}
