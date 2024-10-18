package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	initialMoney = 1000
	dicesThrows  = 3
	minBet       = 1
)

type betOption struct {
	name  string
	sum   []int
	quote int
}

var betOptions = []betOption{{"Aces", []int{3}, 61},
	{"Pequeno", []int{5, 6, 7, 8}, 1},
	{"Grande", []int{14, 15, 16, 17}, 1}}

func main() {
	money := initialMoney
	fmt.Println("*****Banca Francesa*****\nIn jeder Runde koennen Sie einen Teil Ihres Geldes auf eine der folgenden Kombinationen setzen:\n1: Aces, Augensumme: 3 mit einer Gewinnquote von 1 : 61.\n2: Pequeno, Augensumme: 5 6 7 mit einer Gewinnquote von 1 : 1.\n3: Grande, Augensumme: 14 15 16 mit einer Gewinnquote von 1 : 1.")
	for money > 0 {
		fmt.Printf("Money: %d\n", money)
		stake, betIndex := getUserInput(money)
		betPlayer := &(betOptions[betIndex])
		fmt.Printf("%d auf %s gesetzt.\n", stake, betPlayer.name)
		sum := rollDice()
		winningBet := findBetOptionForSum(sum)
		if winningBet != nil {
			fmt.Println(winningBet.name, "!")
		} else {
			fmt.Println("Nothing happens ")
			continue
		}
		win := 0
		switch {
		case betPlayer == winningBet:
			win = stake * betPlayer.quote
			fmt.Printf("You win %d\n", win)
		default:
			win = -stake
			fmt.Printf("You lose %d\n", -win)
		}
		money += win
	}
	fmt.Println("No Money left, game over ")
}

func findBetOptionForSum(sum int) *betOption {
	for i := range betOptions {
		for _, outcome := range betOptions[i].sum {
			if outcome == sum {
				return &betOptions[i]
			}
		}
	}
	return nil
}

func getUserInput(money int) (stake int, betIndex int) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Bet and Combination: ")
		if !scanner.Scan() {
			fmt.Println("Error reading input")
			continue
		}
		input := strings.Fields(scanner.Text())
		if len(input) != 2 {
			fmt.Println("Please enter two numbers separated by a space")
			continue
		}
		var err error
		stake, err = strconv.Atoi(input[0])
		if err != nil {
			fmt.Println("Invalid bet amount")
			continue
		}
		betIndex, err = strconv.Atoi(input[1])
		if err != nil {
			fmt.Println("Invalid combination")
			continue
		}
		betIndex--
		switch {
		case stake > money:
			fmt.Println("Not enough money!")
		case stake < minBet:
			fmt.Println("Bet must be at least", minBet)
		case betIndex < 0 || betIndex >= len(betOptions):
			fmt.Printf("Combination must be between 1 and %d\n", len(betOptions))
		default:
			return
		}
	}
}

func rollDice() (sum int) {
	var dices [dicesThrows]int
	for i := range dices {
		dices[i] = rand.Intn(6) + 1
		sum += dices[i]
	}
	fmt.Printf("Dices: %v with a sum of %d.\n", dices, sum)
	return
}
