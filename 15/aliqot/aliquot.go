package main

import ("fmt";"math/rand"; "time")

func main() {
    rand.Seed(time.Now().UnixNano())
    number := rand.Intn(100)
    var move int
    fmt.Println("*** Aliquot Game ***")
    for {
        fmt.Println(number)
        inputPlayer:
        move = inputPlayer()
        if !isValidMove(move, number) {
            goto inputPlayer
        }
        number -= move
        if gameover(number) {
            fmt.Println("You win !")
            break
        }
        inputComputer:
        move = moveComputer(number)
        if !isValidMove(move, number) {
            goto inputComputer
        }
        fmt.Println("My move: ", move)
        number -= move
        if gameover(number) {
            fmt.Println("You lose !")
            break
        }
    } 
}

func inputPlayer() (movePlayer int) {
    fmt.Print("Your move: ")
    fmt.Scanf("%d", &movePlayer)
    return
}

func moveComputer(number int) int {
    return  rand.Intn(number - 1) + 1 
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
