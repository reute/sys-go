package main

import ("fmt";"math/rand"; "time")

const (
	player = iota
	comp
)

func main() {    
    inputPlayer := func(number int, valid func(int, int) bool) (move int) {
        repeat:
        fmt.Print("Your move: ")
        fmt.Scanf("%d", &move)
        if valid(move, number) {
            return
        } else {
            goto repeat 
        }    
    }    
    inputComputer := func(number int, valid func(int, int) bool) (move int) {
        repeat:
        move = rand.Intn(number - 1) + 1
        if valid(move, number) {
            fmt.Println("My move: ", move)
            return
        } else {
            goto repeat 
        }
    }    
    isValidMove := func(move, number int) bool {
        if move == 0 || number % move != 0 && move != number {
            fmt.Println(move, " is not a proper divisor of ", number)
            return false
        } else {
            return true
        }
    }      
    funcArray := []func(int, func(int, int) bool) int {inputPlayer, inputComputer}    
    rand.Seed(time.Now().UnixNano())
    number := rand.Intn(100)
    var move int
    var turn uint = player
    fmt.Println("*** Aliquot Game ***")  
    for {        
        fmt.Println("Number: ", number)
        move = funcArray[turn](number, isValidMove)
        number -= move 
        if gameover(number) {
            printResult(turn)
            break
        }
        turn += 1; turn = turn % 2           
    }
}

func printResult(turn uint) {
    if turn == comp {
        fmt.Println("You lose !")
    } else if turn == player {
        fmt.Println("You win !")
    }
}

func gameover(number int) bool {
    if number == 1 {
        return true
    } else {
        return false
    }
}