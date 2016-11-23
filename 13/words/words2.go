package main

import ("fmt"; "container/list"; "os";  "bufio"; "strings")

const FILENAME = "cities.txt"

func main() {
    chain := list.New()
    wordList := readFile()
    candidatesLeft := true
    for (candidatesLeft) {
        fmt.Printf("%d Words left\n", wordList.Len())
        word := inputNextWord()  
        if word == "" {
            printHintList(wordList, chain)
        } else {
            e, found := getElement(wordList, word)
            if found {
                if chain.Len() == 0 {
                    wordList, chain = moveBegin(wordList, chain, e)                    
                } else {
                    if fitsBegin(chain, e) {
                        wordList, chain = moveBegin(wordList, chain, e)
                    } else if fitsEnd(chain, e) {
                        wordList, chain = moveEnd(wordList, chain, e)
                    } else {
                        fmt.Println("Word doesn't fit")
                    }
                } 
                printChain(chain)                
                candidatesLeft = check4Candidates(wordList, chain)            
            } else {
                fmt.Println("Word not found in list")
            }
        } 
    }
    fmt.Printf("No further words matching, ending. %d words in Chain", chain.Len()) 
}

func moveBegin(l, c *list.List, e *list.Element) (*list.List, *list.List) {
    c.PushFront(e.Value)     
    l.Remove(e)
    return l, c
}

func moveEnd(l, c *list.List, e *list.Element) (*list.List, *list.List) {
    c.PushBack(e.Value)     
    l.Remove(e)
    return l, c
}

func printChain(c *list.List) {
    fmt.Print(c.Front().Value)
    fmt.Print(" ... ")
    fmt.Print(c.Back().Value)
    fmt.Println()
}

func printHintList(l, c *list.List) {
    for e := l.Front(); e != nil; e = e.Next() {                   
        if fitsBegin(c, e) ||  fitsEnd(c, e) {
            fmt.Printf("%s, ", e.Value) 
        }            
    }  
    fmt.Println()
}

func fitsBegin(c *list.List, e *list.Element) bool {         
    sc := c.Front().Value.(string)    
    se := e.Value.(string)
    se = strings.ToUpper(se)
    se = string(se[len(se)-1])
    return strings.HasPrefix(sc, se)
}

func fitsEnd(c *list.List, e *list.Element) bool {
    sc := c.Back().Value.(string) 
    sc = strings.ToUpper(sc) 
    se := e.Value.(string)
    se = string((se)[0])  
    return strings.HasSuffix(sc, se)
}

func getElement(l *list.List, s string) (*list.Element, bool) {
    var e *list.Element
    for e = l.Front(); e != nil; e = e.Next() {
		if e.Value == s {
            return e, true
        }
	}
    return nil, false   
}

func readFile() *list.List {
    wordList := list.New()
    inFile, _ := os.Open(FILENAME)
    scanner := bufio.NewScanner(inFile)
    scanner.Split(bufio.ScanLines) 
    var item string
    for i := 0; scanner.Scan(); i++ {
        item = scanner.Text()        
        wordList.PushBack(item)
    }
    return wordList
}

func inputNextWord() string {  
    var word string  
    fmt.Print("Next Word: ")
    in := bufio.NewReader(os.Stdin)
	word, _ = in.ReadString('\n')
    word = strings.TrimSpace(word)
    return word
}

func check4Candidates(l, c *list.List) bool {
    for e := l.Front(); e != nil; e = e.Next() {                   
        if fitsBegin(c, e) ||  fitsEnd(c, e) {
            return true 
        }            
    }
    return false    
}
