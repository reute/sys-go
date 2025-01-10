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

type mountainNode struct {
	data *mountain
	next *mountainNode
}

type mountainList struct {
	head *mountainNode
}

func main() {
	unsortedMountains := readFile(NUM_MTNS)
	unsorted := createLinkedList(unsortedMountains)
	sorted := &mountainList{}

	// gameloop
	for unsorted.head != nil {
		fmt.Println("Sorted Mountains:")
		sorted.printMountains()
		fmt.Println("Unsorted Mountains:")
		unsorted.printMountains()

		number_unsorted, number_sorted := inputUser()
		mountainToMove := unsorted.remove(number_unsorted)
		if mountainToMove != nil {
			sorted.insertAt(number_sorted, mountainToMove)
		}
		if !sorted.isSorted() {
			sorted.printMountainsWithHeight()
			fmt.Println(LOSE_MESSAGE)
			return
		}
	}
	fmt.Println(WIN_MESSAGE)
}

func readFile(num_mtns int) []mountain {
	inFile, err := os.Open(FILENAME)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer inFile.Close()

	mountains := make([]mountain, num_mtns)
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for i := 0; scanner.Scan(); i++ {
		parts := strings.Split(scanner.Text(), ":")
		name := parts[0]
		height, _ := strconv.Atoi(parts[1])
		if i < num_mtns {
			mountains[i] = mountain{name, height}
		} else {
			r := rand.Intn(i)
			if r < num_mtns {
				mountains[r] = mountain{name, height}
			}
		}
	}
	return mountains
}

func createLinkedList(mountains []mountain) *mountainList {
	list := &mountainList{}
	for i := range mountains {
		list.append(&mountains[i])
	}
	return list
}

func inputUser() (int, int) {
	fmt.Print(INPUT_PROMPT)
	var from, to int
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}
		input = strings.TrimSpace(input)
		_, err = fmt.Sscanf(input, "%d %d", &from, &to)
		if err == nil {
			break
		}
		fmt.Println("Invalid input. Please enter two numbers.")
	}
	return from, to
}

func (list mountainList) isSorted() bool {
	var tmp, highest int
	current := list.head
	for current != nil {
		tmp = current.data.height
		if tmp < highest {
			return false
		}
		highest = tmp
		current = current.next
	}
	return true
}

func (list *mountainList) append(m *mountain) {
	newNode := &mountainNode{data: m}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *mountainList) remove(index int) *mountain {
	if list.head == nil {
		return nil
	}
	if index == 0 {
		removed := list.head
		list.head = list.head.next
		return removed.data
	}
	current := list.head
	for i := 0; i < index-1 && current.next != nil; i++ {
		current = current.next
	}
	if current.next == nil {
		return nil
	}
	removed := current.next
	current.next = current.next.next
	return removed.data
}

func (list *mountainList) insertAt(index int, m *mountain) {
	newNode := &mountainNode{data: m}
	if index == 0 {
		newNode.next = list.head
		list.head = newNode
		return
	}
	current := list.head
	for i := 0; i < index-1 && current.next != nil; i++ {
		current = current.next
	}
	newNode.next = current.next
	current.next = newNode
}

func (list mountainList) printMountains() {
	current := list.head
	for i := 0; current != nil; i++ {
		fmt.Printf(" %d: %s\n", i, current.data.name)
		current = current.next
	}
}

func (list mountainList) printMountainsWithHeight() {
	current := list.head
	for i := 0; current != nil; i++ {
		fmt.Printf(" %d: %s, Height: %d\n", i, current.data.name, current.data.height)
		current = current.next
	}
}
