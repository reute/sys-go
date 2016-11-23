package main

import ("fmt";     "net/http";    "io/ioutil";    "os"; "strings";  "math/rand";  "time")

const URL = "http://www-cs-faculty.stanford.edu/~uno/sgb-words.txt"
const WORDS_MAX = 9
const (
	comp = iota
	player
)

func main() {
    rand.Seed(time.Now().UnixNano())
    wordList := getWordList()
    var plLists [2][]string
    fmt.Println("*** shack shock ***\nIn this game players alternately select words from this set of words:")
    fmt.Printf("%s\n", wordList)
    fmt.Println("The player wins who is first able to collect 3 occurrences of the same letter.")
    var turn, iWord int
    begin:
    fmt.Print("Do you want to start? Yes=1 No=0 ")
    fmt.Scanf("%d", &turn)
    if turn > 1 || turn < 0 {
        goto begin
    }
    funcArray := []func([]string) int {inputComp, inputPlayer}
    for {  
        fmt.Println("Remaining words: ", wordList)
        iWord = funcArray[turn](wordList)
        plLists[turn] = append(plLists[turn], wordList[iWord])  // add word to playersLists
        wordList = append(wordList[:iWord], wordList[iWord+1:]...)  // remove word from wordList
        if win, char := playerWins(plLists[turn]); win == true {
            fmt.Printf("Got 3 occurrences of %s: %s\n", char, plLists[turn])
            if turn == player {
                fmt.Println("You win !")
            } else if turn == comp {
                fmt.Println("Sorry you lose !")
            }            
            break
        }
        if isDraw(wordList, plLists)  {
            fmt.Println("Game ended with a draw!")
            break
        }
        turn += 1; turn = turn % 2
    }
    fmt.Println("Bye, thanks for playing!")
}

func isDraw(wordList []string, plList [2][]string) bool {
    var joinList []string
    for _, v := range plList {
        joinList = append(v, wordList...)
        if wins, _ := playerWins(joinList); wins == true {
            return false
        }
    }
    return true   
}

func playerWins(plList []string) (bool, string) {
    str := strings.Join(plList, "")
    var c string
    for _, v := range str {
        c = string(v)
        if strings.Count(str, c) >= 3 {
            return true, c
        }
    }
    return false, ""    
}

func inputPlayer(wordList []string) int {
    var word string
    for {
        fmt.Print("You choose: ")
        fmt.Scanf("%s", &word)
        for i, v := range wordList {
            if v == word {
                return i
            }
        }
        fmt.Println("Word not in List !")
    }
}

func inputComp(wordList []string) int {
    r := rand.Intn(len(wordList))
    fmt.Println("The Computer selects", wordList[r] )
    return r
}

func getWordList() []string { 
    var wordList [WORDS_MAX]string   
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
    str := string(contents)
    strSlice := strings.Split(str, "\n")
    var r int   
    for i, v := range strSlice {        
        if i < WORDS_MAX {                
            wordList[i] = v
        } else {
            r = rand.Intn(i)
            if r < WORDS_MAX {
                wordList[r] = v
            }
        }
    }        
    return wordList[:]
}
    
 