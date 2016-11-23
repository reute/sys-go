package main

import ("fmt";  "net/http";    "io/ioutil";    "os"; "strings";  "math/rand";  "time")

const URL = "http://www.mathematik.uni-ulm.de/sai/ws14/soft1/uebungen/blatt05/Ulm.hm"
const MAX_WRONG_GUESSES = 7

func main() {
    guessLeft := MAX_WRONG_GUESSES
    var guesses string
    rand.Seed(time.Now().UnixNano())
    name, details := getPerson()
    name = strings.ToLower(name)
    fmt.Println(name, details)
    hiddenName := getHiddenName(name, guesses)
    fmt.Printf("%d %20s %s\n", guessLeft, guesses, hiddenName)
    for (guessLeft > 0) {        
        c := inputPlayer(guesses)
        if strings.Contains(name, c) {
            guesses += c
            hiddenName = getHiddenName(name, guesses)
        } else {
            guessLeft -= 1 
        }
        fmt.Printf("%d %-20s %s\n", guessLeft, guesses, hiddenName)
        if playerWins(hiddenName) {
            fmt.Println("Congratulations, you got", guessLeft, "points")
            os.Exit(0)
        }
    }
    fmt.Println("You lose. The correct answer was", name)   
}

func playerWins(hiddenName string) bool {
     if (strings.Contains(hiddenName, "_")) {
         return false
     }  else {
         return true
     }
}

func inputPlayer(guesses string) (c string) {
    input:
    fmt.Print("Input Char: ")
    fmt.Scanf("%s", &c) 
    if isValidInput(c, guesses) {
        return
    } else {
        goto input
    }
}

func isValidInput(c, guesses string) bool {
    if len(c) > 1 {
        fmt.Println("Only one char please ") 
        return false       
    } else if (strings.Contains(guesses, c)) {
        fmt.Println("Character already taken, please choose another") 
        return false
    }
    return true
}

func getHiddenName(name, guesses string) string {
    var hiddenString string
    var c string
    for _, v := range name {
        c = string(v)
        if c != " " && !strings.Contains(guesses, c)   {  
            hiddenString += "_"
        } else {
            hiddenString += c
        }
    }
    return hiddenString
}

func getPerson() (name, details string) {  
    response, err := http.Get(URL)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } 
    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }
    oneLine := string(contents)
    lines := strings.Split(oneLine, "\n")
    r := rand.Intn(len(lines))
    line := strings.Split(lines[r], "\t")
    name = line[1]
    details = line[0]       
    return
}
    
 