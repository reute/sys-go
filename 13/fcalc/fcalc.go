package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []float32

func main() {
	fmt.Print("Function: ")
	// Reverse Polish Notation is e.g. "6 3 + 1 -"
	input := inputFunction()
	result, err := calculate(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Result: %v\n", result)
}

func inputFunction() string {
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	return str
}

func calculate(expression string) (float32, error) {
	expr_slice := strings.Fields(expression)
	var stack stack
	for _, token := range expr_slice {
		switch token {
		case "+":
			op2, err := stack.Pop()
			if err != nil {
				return 0, fmt.Errorf("stack is empty")
			}
			op1, err := stack.Pop()
			if err != nil {
				return 0, fmt.Errorf("stack is empty")
			}
			stack.Put(op1 + op2)
		case "-":
			op2, err := stack.Pop()
			if err != nil {
				return 0, fmt.Errorf("stack is empty")
			}
			op1, err := stack.Pop()
			if err != nil {
				return 0, fmt.Errorf("stack is empty")
			}
			stack.Put(op1 - op2)
		case "*":
			op2, err := stack.Pop()
			if err != nil {
				return 0, fmt.Errorf("stack is empty")
			}
			op1, err := stack.Pop()
			if err != nil {
				return 0, fmt.Errorf("stack is empty")
			}
			stack.Put(op1 * op2)
		case "/":
			op2, err := stack.Pop()
			if err != nil {
				return 0, fmt.Errorf("stack is empty")
			}
			op1, err := stack.Pop()
			if err != nil {
				return 0, fmt.Errorf("stack is empty")
			}
			stack.Put(op1 / op2)
		default:
			num, err := strconv.ParseFloat(token, 32)
			if err != nil {
				return 0, fmt.Errorf("invalid token: %s", token)
			}
			stack.Put(float32(num))
		}
	}
	result, err := stack.Pop()
	return result, err
}

func (s stack) Peek() float32 {
	return s[len(s)-1]
}

func (s *stack) Put(f float32) {
	*s = append(*s, f)
}

func (s *stack) Pop() (float32, error) {
	size := len(*s)
	if size == 0 {
		return 0, errors.New("stack is empty")
	}
	f := (*s).Peek()
	*s = (*s)[:size-1]
	return f, nil
}
