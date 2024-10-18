package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const MAX_WRONG_GUESSES = 7

var persons = []string{
	"Albert Einstein\tTheoretical Physicist",
	"Marie Curie\tPhysicist and Chemist",
	"Isaac Newton\tPhysicist and Mathematician",
	"Charles Darwin\tNaturalist and Biologist",
	"Nikola Tesla\tInventor and Electrical Engineer",
	"Galileo Galilei\tAstronomer and Physicist",
	"Stephen Hawking\tTheoretical Physicist",
	"Ada Lovelace\tMathematician and Computer Programmer",
	"Leonardo da Vinci\tPolymath and Artist",
	"Rosalind Franklin\tChemist and X-ray Crystallographer",
}

func main() {
	guessLeft := MAX_WRONG_GUESSES
	var guesses string
	name, details := getPerson()
	name = strings.ToLower(name)
	fmt.Println(name, details)
	hiddenName := getHiddenName(name, guesses)
	fmt.Printf("%d %20s %s\n", guessLeft, guesses, hiddenName)
	for guessLeft > 0 {
		c := inputPlayer(guesses)
		guesses += c
		if strings.Contains(name, c) {
			hiddenName = getHiddenName(name, guesses)
		} else {
			guessLeft--
		}
		fmt.Printf("%d %-20s %s\n", guessLeft, guesses, hiddenName)
		if playerWins(hiddenName) {
			fmt.Println("Congratulations, you got", guessLeft, "points")
			os.Exit(0)
		}
	}
	fmt.Println("You lose. The correct answer was", name)
}

func getPerson() (name, details string) {
	line := strings.Split(persons[rand.Intn(len(persons))], "\t")
	return line[0], line[1]
}

func getHiddenName(name, guesses string) string {
	var hiddenString strings.Builder
	for _, v := range name {
		c := string(v)
		if c != " " && !strings.Contains(guesses, c) {
			hiddenString.WriteRune('_')
		} else {
			hiddenString.WriteString(c)
		}
	}
	return hiddenString.String()
}

func inputPlayer(guesses string) (c string) {
	for {
		fmt.Print("Input Char: ")
		var c string
		fmt.Scanf("%s", &c)
		if isValidInput(c, guesses) {
			return c
		}
	}
}

func isValidInput(c, guesses string) bool {
	if len(c) > 1 {
		fmt.Println("Only one char please ")
		return false
	} else if strings.Contains(guesses, c) {
		fmt.Println("Character already taken, please choose another")
		return false
	}
	return true
}

func playerWins(hiddenName string) bool {
	return !strings.Contains(hiddenName, "_")
}
