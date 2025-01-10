package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const (
	begin = iota
	end
	nofit
	running
	noCandidatesLeft
	noCitiesLeft
)
const filename = "cities.txt"

type Node struct {
	city string
	next *Node
}

func main() {
	unsorted, err := readFromFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		return
	}
	var sorted *Node
	gameStatus := checkGameStatus(unsorted, sorted)
	for gameStatus == running {
		fmt.Printf("%d Words left\n", unsorted.len())
		inputCity := getInputCity()
		if inputCity == "" {
			printHintList(findCandidates(unsorted, sorted))
		} else {
			cityNode := searchCity(inputCity, unsorted)
			if cityNode != nil {
				position := checkFit(inputCity, sorted)
				if position != nofit {
					insertIntoSorted(&unsorted, &sorted, inputCity, position)
					sorted.print()
					gameStatus = checkGameStatus(unsorted, sorted)
				} else {
					fmt.Println("Word doesn't fit")
				}
			} else {
				fmt.Println("Word not found in list")
			}
		}
	}
	switch gameStatus {
	case noCandidatesLeft:
		fmt.Printf("No further candidates available, ending. %d words in Chain", sorted.len())
	case noCitiesLeft:
		fmt.Printf("No further cities available, ending. %d words in Chain", sorted.len())
	}
}

func insertIntoSorted(unsorted **Node, sorted **Node, city string, position int) {
	// Find the node in the unsorted list
	var prev *Node
	curr := *unsorted
	for curr != nil && curr.city != city {
		prev = curr
		curr = curr.next
	}
	// If the node is not found, return
	if curr == nil {
		fmt.Println("City not found in the unsorted list")
		return
	}
	// Remove the node from the unsorted list
	if prev == nil {
		*unsorted = curr.next
	} else {
		prev.next = curr.next
	}
	// Insert the node into the sorted list
	if *sorted == nil || position == begin {
		// Insert at the beginning
		curr.next = *sorted
		*sorted = curr
	} else {
		// Insert at the end
		sortedCurr := *sorted
		for sortedCurr.next != nil {
			sortedCurr = sortedCurr.next
		}
		sortedCurr.next = curr
		curr.next = nil
	}
}

func checkFit(city string, sorted *Node) int {
	if sorted == nil {
		return begin
	}
	if fitsFirstChar(city, sorted.city) {
		return begin
	}
	if fitsLastChar(city, sorted.getLastNode().city) {
		return end
	}
	return nofit
}

func fitsFirstChar(cityNew, citySorted string) bool {
	return strings.EqualFold(string(cityNew[len(cityNew)-1]), string(citySorted[0]))
}

func fitsLastChar(cityNew, citySorted string) bool {
	return strings.EqualFold(string(cityNew[0]), string(citySorted[len(citySorted)-1]))
}

func readFromFile(filename string) (unsorted *Node, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lastNode *Node
	for scanner.Scan() {
		cityNode := &Node{
			city: scanner.Text(),
		}
		if unsorted == nil {
			unsorted = cityNode
		} else {
			lastNode.next = cityNode
		}
		lastNode = cityNode
	}
	return
}

func findCandidates(unsorted *Node, sorted *Node) (candidates []string) {
	if unsorted == nil {
		return
	}
	current := unsorted
	for current != nil {
		city := current.city
		position := checkFit(city, sorted)
		if position != nofit {
			candidates = append(candidates, city)
		}
		current = current.next
	}
	return
}

func (list *Node) getLastNode() *Node {
	if list == nil {
		return list
	}
	current := list
	for current.next != nil {
		current = current.next
	}
	return current
}

func (list *Node) len() int {
	length := 0
	for current := list; current != nil; current = current.next {
		length++
	}
	return length
}

func (list *Node) print() {
	fmt.Print(list.city)
	fmt.Print(" ... ")
	fmt.Print(list.getLastNode().city)
	fmt.Println()
}

func checkGameStatus(unsorted, sorted *Node) int {
	if unsorted.len() == 0 {
		return noCitiesLeft
	}
	if sorted.len() == 0 {
		return running
	}
	candidates := findCandidates(unsorted, sorted)
	if len(candidates) != 0 {
		return running
	} else {
		return noCandidatesLeft
	}
}

func getInputCity() (inputCity string) {
	fmt.Print("Next City: ")
	tmp := bufio.NewReader(os.Stdin)
	inputCity, _ = tmp.ReadString('\n')
	inputCity = strings.TrimSpace(inputCity)
	return
}

func printHintList(candidates []string) {
	if len(candidates) == 0 {
		fmt.Println("No candidates available.")
		return
	}
	sort.Strings(candidates)
	fmt.Println("Available candidates:")
	for i, city := range candidates {
		if i > 0 && i%5 == 0 {
			fmt.Println()
		}
		fmt.Printf("%-20s", city)
	}
	fmt.Println()
}

func searchCity(inputCity string, unsorted *Node) *Node {
	current := unsorted
	for current != nil {
		if strings.EqualFold(inputCity, current.city) {
			return current
		}
		current = current.next
	}
	return nil
}
