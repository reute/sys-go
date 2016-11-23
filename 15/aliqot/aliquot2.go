package main

import ("fmt";"math/rand"; "time")

const (
	player = iota
	comp
)

func getInput(number int, turn uint) (move int)  {    
    if turn == player {        
        fmt.Print("Your move: ")
        fmt.Scanf("%d", &move)
    } else if turn == comp {
        move = rand.Intn(number - 1) + 1
        fmt.Println("My move: ", move) 
    }
    return
}

func main() {
    rand.Seed(time.Now().UnixNano())
    number := rand.Intn(100)
    var move int
    var turn uint = player
    fmt.Println("*** Aliquot Game ***")  
    for {        
        fmt.Println(number)        
        input:
        move = getInput(number, turn)
        if !isValidMove(move, number) {
            goto input
        }
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

func isValidMove(move, number int) bool {
    if number % move != 0 && move != number {
        fmt.Println(move, " is not a proper divisor of ", number)
        return false
    } else {
        return true
    }
}

func gameover(number int) bool {
    if number == 1 {
        return true
    } else {
        return false
    }
}