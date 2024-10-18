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
	NUM_MTNS     = 8
	FILENAME     = "berge"
	WIN_MESSAGE  = "You have won!"
	LOSE_MESSAGE = "Sorry, the mountains are no longer sorted."
	INPUT_PROMPT = "What is to be inserted where? "
)

type mountain struct {
	name   string
	height int
}

type data struct {
	unsorted []*mountain
	sorted   []*mountain
}

func main() {
	mountains := readFile()
	game := newGame(mountains)

	// gameloop
	for len(game.unsorted) > 0 {
		fmt.Println("Sorted Mountains:")
		printMountains(game.sorted)
		fmt.Println("Still to be sorted:")
		printMountains(game.unsorted)

		from, to := inputUser()
		game.move(from, to)

		if !isSorted(game.sorted) {
			break
		}
	}
	if len(game.unsorted) == 0 {
		fmt.Println("You have  won !!")
	} else {
		fmt.Println("Sorry, then it is no longer sorted:")
	}
	for _, element := range game.sorted {
		fmt.Printf(" %5d: %s \n", element.height, element.name)
	}
	fmt.Println("Bye!")
	fmt.Printf("You got %d Points.\n", len(game.sorted)-1)
}

func readFile() (mountains [NUM_MTNS]mountain) {
	inFile, err := os.Open(FILENAME)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	var tmp []string
	var name string
	var height, r int
	for i := 0; scanner.Scan(); i++ {
		tmp = strings.Split(scanner.Text(), ":")
		name = tmp[0]
		height, _ = strconv.Atoi(tmp[1])
		if i < NUM_MTNS {
			mountains[i] = mountain{name, height}
		} else {
			r = rand.Intn(i)
			if r < NUM_MTNS {
				mountains[r] = mountain{name, height}
			}
		}
	}
	return
}

func newGame(allMountains [NUM_MTNS]mountain) *data {
	game := &data{
		unsorted: make([]*mountain, len(allMountains)),
		sorted:   make([]*mountain, 0, len(allMountains)),
	}
	for i := range allMountains {
		game.unsorted[i] = &allMountains[i]
	}
	return game
}

func inputUser() (int, int) {
	fmt.Print(INPUT_PROMPT)
	var from, to int
	_, err := fmt.Scanf("%d %d", &from, &to)
	if err != nil {
		fmt.Println("Invalid input. Please enter two numbers.")
		return inputUser()
	}
	return from, to
}

func isSorted(mountains []*mountain) bool {
	var tmp, hightest int
	for _, element := range mountains {
		tmp = element.height
		if tmp < hightest {
			return false
		}
		hightest = tmp
	}
	return true
}

func (m *data) move(from int, to int) {
	m.sorted = append(m.sorted, &mountain{})
	copy(m.sorted[to+1:], m.sorted[to:])
	m.sorted[to] = m.unsorted[from]
	m.unsorted = append(m.unsorted[:from], m.unsorted[from+1:]...)
}

func printMountains(mountains []*mountain) {
	for i, element := range mountains {
		fmt.Printf(" %d: %s\n", i, element.name)
	}
}
