package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type bet struct {
	name  string
	sum   []int
	quote int
}

func main() {
	rand.Seed(4)
	money := 1000

	bets := [3]bet{bet{"Aces", []int{3}, 61},
		bet{"Pequeno", []int{5, 6, 7, 8}, 1},
		bet{"Grande", []int{14, 15, 16, 17}, 1}}
	var win, stake, betIndex, sum int
	var betDice, betPlayer *bet
	var dice [3]int

	fmt.Println("*****Banca Francesa*****\nIn jeder Runde koennen Sie einen Teil Ihres Geldes auf eine der folgenden Kombinationen setzen:\n1: Aces, Augensumme: 3 mit einer Gewinnquote von 1 : 61.\n2: Pequeno, Augensumme: 5 6 7 mit einer Gewinnquote von 1 : 1.\n3: Grande, Augensumme: 14 15 16 mit einer Gewinnquote von 1 : 1.")

	for money > 0 {
		fmt.Printf("Money: %d\n", money)
		stake, betIndex = inputUser(money)
		betPlayer = &(bets[betIndex])
		fmt.Printf("%d auf %s gesetzt.\n", stake, betPlayer.name)
		for {
			dice = rollDice()
			sum = dice[0] + dice[1] + dice[2]
			fmt.Printf("Dice: %d %d %d with a sum of %d.\n", dice[0], dice[1], dice[2], sum)
			betDice = findSumInBet(&bets, sum)
			if betDice != nil {
				break
			}
			fmt.Println("Nothing happens ")
		}
		fmt.Println(betDice.name, "!")
		if betPlayer == betDice {
			win = stake * betPlayer.quote
			fmt.Println("You win ", win)
		} else {
			win = stake * -1
			fmt.Println("You lose", win*-1)
		}
		money += win
		win = 0
	}
	fmt.Println("No Money left, game over ")
}

func findSumInBet(bets *[3]bet, sumdice int) *bet {
    var bet *bet
	for i := 0; i < len(bets); i++ {
		bet = &(bets[i])
		for _, sumbet := range bet.sum {
			if sumbet == sumdice {
				return bet
			}
		}
	}
	return nil
}

func inputUser(money int) (stake int, combi int) {    
	for {
		fmt.Print("Bet and Combination: ")
		in := bufio.NewReader(os.Stdin)
		str, _ := in.ReadString('\n')
		strSlice := strings.Fields(str)
		stake, _ = strconv.Atoi(strSlice[0])
		combi, _ = strconv.Atoi(strSlice[1])
		combi -= 1
		if stake > money {
			fmt.Println("Not enough Money ! ")
		} else if money < 1 {
			fmt.Println("Bet must be bigger than 0 ")
		} else if combi > 2 || combi < 0 {
			fmt.Println("Combination must be between 1 and 3")
		} else {
			return 
		}
	}
}

func rollDice() (dice [3]int) {
	for i := 0; i < len(dice); i++ {
		dice[i] = rand.Intn(6)
	}
    return
}
