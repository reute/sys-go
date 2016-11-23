package main

import (
	"fmt"
	"bufio" 
	"os"
	"strings"
)

const clear string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "
const ANZAHL_ZEICHEN int = len(clear)

func main() {    
	fmt.Println("*** ENIGMA ***")
	text := inputText()
    mode := inputMode()	
    var result string    
	switch mode {
		case 1: 
            walzenstellung := inputWalzenstellung()
            result = encrypt(text, walzenstellung)
            fmt.Printf("\nVerschlüsselter Text:\n%s\n", result)
		case 2: 
            walzenstellung := inputWalzenstellung()
            result = decrypt(text, walzenstellung)
            fmt.Printf("\nEntschlüsselter Text:\n%s\n", result)
        case 3: 
            word := inputClearWord()
            result = crack(text, word)
            if (result == "") {
                fmt.Println("Suchwort nicht gefunden !")
            } else {
                fmt.Printf("Gecrackter Text:\n%s\n", result);
            }
    }          
}

func getWalze(walzenstellung int) string {
    walze := "EKLMF6GDQVZ0TO8Y XUSP2IB4CJ5AR197W3NH"
    for i := 0; i < walzenstellung; i++ {
        walze = walze[1:len(walze)] + string(walze[0])
	}
    return walze
}

// with slices
func encrypt(text string, walzenstellung int) string {
    walze := getWalze(walzenstellung) 
    result := make([]string, len(text))
    for it, runechar := range text {
      	if !strings.ContainsRune(clear, runechar) {			
			runechar = ' '
		}
        ic := strings.IndexRune(clear, runechar) 
        result[it] = string(walze[ic])
        walze = walze[1:len(walze)] + string(walze[0])
    } 
    return strings.Join(result, "")
 }

//without slices 
func decrypt(text string, walzenstellung int) string {
    walze := getWalze(walzenstellung)
    var result string
	for _, runechar := range text {
        i := strings.IndexRune(walze, runechar)
        result += clear[i:i+1]
        walze = walze[1:len(walze)] + walze[0:1]
	}	
    return result
}

func crack(text string, word string) string {
    var result string
    for i := 0; i < ANZAHL_ZEICHEN; i++ {        
        result = decrypt(text, i)
        if (strings.Contains(text, word)) {
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
    fmt.Printf("Bitte Walzenstellung eingeben (0 - %d): ", ANZAHL_ZEICHEN)
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