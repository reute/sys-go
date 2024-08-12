package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const MAX_MOUNTAINS = 8
const FILENAME = "berge"

type mountain struct {
	name   string
	height int
}

type mountains struct {
	unsorted []*mountain
	sorted   []*mountain
}

func main() {
	var from, to int
	all_mountains := readFile()
	game := mountains{
		unsorted: make([]*mountain, len(all_mountains)),
		sorted:   make([]*mountain, 0),
	}
	for i := 0; i < len(all_mountains); i++ {
		game.unsorted[i] = &all_mountains[i]
	}
	// gameloop
	for isSorted(game.sorted) && len(game.unsorted) != 0 {
		fmt.Println("Sorted Mountains:")
		printMountains(game.sorted)
		fmt.Println("Still to be sorted:")
		printMountains(game.unsorted)
		from, to = inputUser()
		game.move(from, to)
	}
	if len(game.unsorted) == 0 {
		fmt.Println("You have  won !!")
	} else {
		fmt.Println("Sorry, then it is no longer sorted:")
	}
	for _, element := range game.sorted {
		fmt.Printf(" %5d: %s \n", (*element).height, (*element).name)
	}
	fmt.Println("Bye!")
	fmt.Printf("You got %d Points.\n", len(game.sorted)-1)
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

func inputUser() (int, int) {
	fmt.Println("What is to be inserted where? ")
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	strSlice := strings.Fields(str)
	from, _ := strconv.Atoi(strSlice[0])
	to, _ := strconv.Atoi(strSlice[1])
	return from, to
}

func readFile() (mountains [MAX_MOUNTAINS]mountain) {
	inFile, _ := os.Open(FILENAME)
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
		if i < MAX_MOUNTAINS {
			mountains[i] = mountain{name, height}
		} else {
			r = rand.Intn(i)
			if r < MAX_MOUNTAINS {
				mountains[r] = mountain{name, height}
			}
		}
	}
	return
}

func (m *mountains) move(from int, to int) {
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
