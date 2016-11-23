package main

import ("fmt"; "bufio"; "os"; "strings")

const CLEAR = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "
const WALZE = "EKLMF6GDQVZ0TO8Y XUSP2IB4CJ5AR197W3NH"
const ANZAHL_ZEICHEN = len(CLEAR)

type enigma struct { 
    text, walze string       
}

func main() {    
	fmt.Println("*** ENIGMA ***")    
	text := inputText()
    mode := inputMode()
    switch mode {
		case 1:  
            walzenstellung := inputWalzenstellung()
            e := enigma{text, WALZE}
            e.initWalze(walzenstellung)           
            result := e.encrypt()
            fmt.Printf("\nVerschlüsselter Text:\n%s\n", result)
		case 2: 
            walzenstellung := inputWalzenstellung()
            e := enigma{text, WALZE}
            e.initWalze(walzenstellung) 
            result := e.decrypt()
            fmt.Printf("\nEntschlüsselter Text:\n%s\n", result)
        case 3: 
            word := inputClearWord()
            result := crack(text, word)
            if (result == "") {
                fmt.Println("Suchwort nicht gefunden !")
            } else {
                fmt.Printf("Gecrackter Text:\n%s\n", result);
            }
    }          
}

func (e *enigma) initWalze(walzenstellung int) { 
    for i := 0; i < walzenstellung; i++ {
        e.walze = e.walze[1:len(e.walze)] + string(e.walze[0])
	}
}

func (e *enigma) encrypt() string {     
    result := make([]string, len(e.text))
    for it, runechar := range e.text {
      	if !strings.ContainsRune(CLEAR, runechar) {			
			runechar = ' '
		}
        ic := strings.IndexRune(CLEAR, runechar) 
        result[it] = string(e.walze[ic])
        e.walze = e.walze[1:len(e.walze)] + string(e.walze[0])
    } 
    return strings.Join(result, "")
}

// ohne slice
func (e *enigma) decrypt() string {
    var result string
	for _, runechar := range e.text {
        i := strings.IndexRune(e.walze, runechar)
        result += string(CLEAR[i])
        e.walze = e.walze[1:len(e.walze)] + string(e.walze[0])
	}	
    return result
}

func crack(text string, word string) string {
    var result string
    var e enigma
    for walzenstellung := 0; walzenstellung < ANZAHL_ZEICHEN; walzenstellung++ {
        e = enigma{text, WALZE}
        e.initWalze(walzenstellung)        
        result = e.decrypt()
        if (strings.Contains(result, word)) {
            return result
        }			
	}
    return ""
}

func inputClearWord() string {
    var word string
    fmt.Println("Bitte Suchwort eingeben: ")
    fmt.Scanf("%s", &word)
    return word
}

func inputText() string {
   	fmt.Println("Bitte Text eingeben:")
    in := bufio.NewReader(os.Stdin)
	text, _ := in.ReadString('\n')
	text = strings.TrimSpace(text)
    text = strings.ToUpper(text)
    return text
}

func inputWalzenstellung() int {
    var walzenstellung int
    fmt.Printf("Bitte Walzenstellung eingeben (0 - %d): ", 37)
	for {
        fmt.Scanf("%d", &walzenstellung)
        if walzenstellung >= 0 && walzenstellung <= 37 {
            return walzenstellung
        } else {
            fmt.Print("Bitte Wert zwischen 0 und 37 eingeben: ")
        }  
    }  
}

func inputMode() int {
    var mode int
    fmt.Println("1 - Verschlüsseln\n2 - Entschlüsseln\n3 - Text Knacken")
    for {
       	fmt.Scanf("%d", &mode)
        if mode >= 1 && mode <= 3 {
            return mode
        } else {
            fmt.Print("Bitte Wert zwischen 1 und 3 eingeben")
        }  
    } 
}