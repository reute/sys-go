package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	stdin, _ := io.ReadAll(reader)
	text := string(stdin)
	// fmt.Println("Something on STDIN: " + text)
	re := regexp.MustCompile(`http://[^"]*`)
	result := re.FindAllString(text, -1)
	for i, val := range result {
		fmt.Printf(" %d: %s\n", i, val)
	}
}
