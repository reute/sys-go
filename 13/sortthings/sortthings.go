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

type mountainNode struct {
	name   string
	height int
	next   *mountainNode
}

type mountainList struct {
	head *mountainNode
}

func main() {
	unsorted := readFile(NUM_MTNS)
	sorted := &mountainList{}

	for unsorted.head != nil {
		fmt.Println("Sorted Mountains:")
		sorted.printMountains()
		fmt.Println("Unsorted Mountains:")
		unsorted.printMountains()

		number_unsorted, number_sorted := inputUser()
		node := unsorted.remove(number_unsorted)
		if node != nil {
			sorted.insertAt(number_sorted, node)
		}
		if !sorted.isSorted() {
			sorted.printMountainsWithHeight()
			fmt.Println(LOSE_MESSAGE)
			return
		}
	}
	fmt.Println(WIN_MESSAGE)
}

func readFile(num_mtns int) *mountainList {
	inFile, err := os.Open(FILENAME)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer inFile.Close()
	list := &mountainList{}
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for i := 0; scanner.Scan(); i++ {
		parts := strings.Split(scanner.Text(), ":")
		name := parts[0]
		height, _ := strconv.Atoi(parts[1])
		if i < num_mtns {
			list.append(&mountainNode{name, height, nil})
		} else {
			r := rand.Intn(i)
			if r < num_mtns {
				list.changeAt(r, name, height)
			}
		}
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
		tmp = current.height
		if tmp < highest {
			return false
		}
		highest = tmp
		current = current.next
	}
	return true
}

func (list *mountainList) append(node *mountainNode) {
	if list.head == nil {
		list.head = node
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
}

func (list *mountainList) remove(index int) *mountainNode {
	if list.head == nil {
		return nil
	}
	if index == 0 {
		removed := list.head
		list.head = list.head.next
		return removed
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
	return removed
}

func (list *mountainList) insertAt(index int, node *mountainNode) {
	currentNode := list.head
	if index == 0 {
		list.head = node
		node.next = currentNode
		return
	}
	if currentNode != nil {
		for i := 0; i < index-1 && currentNode.next != nil; i++ {
			currentNode = currentNode.next
		}
	}
	node.next = currentNode.next
	currentNode.next = node
}

func (list *mountainList) changeAt(index int, name string, height int) {
	node := list.head
	for i := 0; i < index-1 && node.next != nil; i++ {
		node = node.next
	}
	node.name = name
	node.height = height
}

func (list mountainList) printMountains() {
	current := list.head
	for i := 0; current != nil; i++ {
		fmt.Printf(" %d: %s\n", i, current.name)
		current = current.next
	}
}

func (list mountainList) printMountainsWithHeight() {
	current := list.head
	for i := 0; current != nil; i++ {
		fmt.Printf(" %d: %s, Height: %d\n", i, current.name, current.height)
		current = current.next
	}
}
