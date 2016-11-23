package main

import ("fmt"; "container/list"; "os";  "bufio"; "strings")

const FILENAME = "cities.txt"

func main() {
    chain := list.New()
    wordList := readFile()
    hintList := wordList
    candidatesLeft := true
    for (candidatesLeft) {
        fmt.Printf("%d Words left\n", wordList.Len())
        word := inputNextWord()  
        if word == "" {
            printHintList(hintList)
        } else {
            e, found := containsString(wordList, word)
            if found {
                if chain.Len() == 0 {
                    wordList, chain = moveBegin(wordList, chain, e)                    
                } else {
                    if fitsBegin(chain, word) {
                        wordList, chain = moveBegin(wordList, chain, e)
                    } else if fitsEnd(chain, word) {
                        wordList, chain = moveEnd(wordList, chain, e)
                    } else {
                        fmt.Println("Word doesn't fit")
                    }
                } 
                printChain(chain)
                hintList = getHintList(wordList, chain)
                candidatesLeft = hintList.Len() > 0             
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

func printHintList(h *list.List) {
    for e := h.Front(); e != nil; e = e.Next() {                        
        fmt.Printf("%s, ", e.Value)        
    }
    fmt.Println()
}

func fitsBegin(c *list.List, s string) bool {         
    v, _ := c.Front().Value.(string)    
    s = strings.ToUpper(s)
    s = string(s[len(s)-1])
    return strings.HasPrefix(v, s)
}

func fitsEnd(c *list.List, s string) bool {
    v, _ := c.Back().Value.(string) 
    v = strings.ToUpper(v) 
    s = string((s)[0])  
    return strings.HasSuffix(v, s)
}

func containsString(l *list.List, s string) (*list.Element, bool) {
    var e *list.Element
    for e = l.Front(); e != nil; e = e.Next() {
		if e.Value == s {
            return e, true
        }
	}
    return e, false   
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

func getHintList(l, c *list.List) *list.List {
    hintList := list.New()
    for e := l.Front(); e != nil; e = e.Next() {                   
        if fitsBegin(c, e.Value.(string)) ||  fitsEnd(c, e.Value.(string)) {
            hintList.PushBack(e.Value) 
        }            
    }
    return hintList    
}
