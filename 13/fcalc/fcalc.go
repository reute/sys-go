package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []float32

func (s stack) Peek() float32 {
	return s[len(s)-1]
}

func (s *stack) Put(f float32) {
	*s = append(*s, f)
}

func (s *stack) Pop() float32 {
	f := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return f
}

func main() {
	fmt.Print("Function: ")
	function := inputFunction()
	//function := "6 3 + 1 -"
	result := calculate(strings.Fields(function))
	fmt.Printf("Result : %f\n", result)
}

func calculate(function []string) float32 {
	var stack stack
	for _, token := range function {
		switch token {
		case "+":
			op2 := stack.Pop()
			op1 := stack.Pop()
			stack.Put(op1 + op2)
		case "-":
			op2 := stack.Pop()
			op1 := stack.Pop()
			stack.Put(op1 - op2)
		case "*":
			op2 := stack.Pop()
			op1 := stack.Pop()
			stack.Put(op1 * op2)
		case "/":
			op2 := stack.Pop()
			op1 := stack.Pop()
			stack.Put(op1 / op2)
		default:
			num, err := strconv.ParseFloat(token, 32)
			if err == nil {
				stack.Put(float32(num))
			} else {
				fmt.Printf("Invalid token: %s\n", token)
			}
		}
	}
	return stack.Pop()
}

func inputFunction() string {
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	return str
}
