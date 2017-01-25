package main

import ("fmt"; "bufio"; "os"; "strings")

const clear = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "
const defaultWalze = "EKLMF6GDQVZ0TO8Y XUSP2IB4CJ5AR197W3NH"
const sizeWalze = len(clear)

type enigma struct { 
    text, walze string       
}

const (
    modeEncrypt = iota
    modeDecrypt
    modeCrack
)

func main() {    
	fmt.Println("*** ENIGMA ***")    
	text := inputText()
    mode := inputMode()
    switch mode {
		case modeEncrypt:  
            walze := getWalze()
            e := enigma{text, walze}           
            result := e.encrypt()
            fmt.Printf("\nVerschlüsselter Text:\n%s\n", result)
		case modeDecrypt: 
            walze :=getWalze()
            e := enigma{text, walze}          
            result := e.decrypt()
            fmt.Printf("\nEntschlüsselter Text:\n%s\n", result)
        case modeCrack: 
            word := inputClearWord()
            result := crack(text, word)
            if (result == "") {
                fmt.Println("Suchwort nicht gefunden !")
            } else {
                fmt.Printf("Gecrackter Text:\n%s\n", result);
            }
    }     
}

func (e *enigma) encrypt() string {     
    result := make([]string, len(e.text))
    for it, runechar := range e.text {
      	if !strings.ContainsRune(clear, runechar) {			
			runechar = ' '
		}
        ic := strings.IndexRune(clear, runechar) 
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
        result += string(clear[i])
        e.walze = e.walze[1:len(e.walze)] + string(e.walze[0])
	}	
    return result
}

func crack(text string, word string) (result string) {   
    var e enigma
    walze := defaultWalze
    for walzenstellung := 0; walzenstellung < sizeWalze; walzenstellung++ {        
        e = enigma{text, walze}      
        result = e.decrypt()
        if (strings.Contains(result, word)) {
            return
        }	
        walze = walze[1:len(walze)] + string(walze[0])		
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

func getWalze() (walze string) {
    var walzenstellung int
    walze = defaultWalze
    fmt.Printf("Bitte Walzenstellung eingeben (0 - %d): ", 37)
	for {
        fmt.Scanf("%d", &walzenstellung)
        if walzenstellung >= 0 && walzenstellung <= 37 {
            break
        } else {
            fmt.Print("Bitte Wert zwischen 0 und 37 eingeben: ")
        }  
    }
    for i := 0; i < walzenstellung; i++ {
        walze = walze[1:len(walze)] + string(walze[0])
	} 
    return 
}

func inputMode() int {
    var mode int
    fmt.Println("0 - Verschlüsseln\n1 - Entschlüsseln\n2 - Text Knacken")
    for {
       	fmt.Scanf("%d", &mode)
        if mode >= 0 && mode <= 2 {
            return mode
        } else {
            fmt.Print("Bitte Wert zwischen 0 und 2 eingeben")
        }  
    } 
}