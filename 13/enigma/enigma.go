package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Use constants for magic numbers and strings
const (
	clear        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "
	defaultWalze = "EKLMF6GDQVZ0TO8Y XUSP2IB4CJ5AR197W3NH"
	sizeWalze    = len(clear)
)

const (
	modeEncrypt = iota
	modeDecrypt
	modeCrack
)

type enigma struct {
	text, walze string
}

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
		walze := getWalze()
		e := enigma{text, walze}
		result := e.decrypt()
		fmt.Printf("\nEntschlüsselter Text:\n%s\n", result)
	case modeCrack:
		word := inputClearWord()
		result := crack(text, word)
		if result == "" {
			fmt.Println("Suchwort nicht gefunden !")
		} else {
			fmt.Printf("Gecrackter Text:\n%s\n", result)
		}
	}
}

func (e *enigma) encrypt() string {
	var result string
	for _, runechar := range e.text {
		if !strings.ContainsRune(clear, runechar) {
			runechar = ' '
		}
		i := strings.IndexRune(clear, runechar)
		result += string(e.walze[i])
		e.walze = e.walze[1:] + e.walze[:1]
	}
	return result
}

func (e *enigma) decrypt() string {
	var result string
	for _, runechar := range e.text {
		i := strings.IndexRune(e.walze, runechar)
		result += string(clear[i])
		e.walze = e.walze[1:] + e.walze[:1]
	}
	return result
}

func crack(text string, word string) (result string) {
	var e enigma
	walze := defaultWalze
	for walzenstellung := 0; walzenstellung < sizeWalze; walzenstellung++ {
		e = enigma{text, walze}
		result = e.decrypt()
		if strings.Contains(result, word) {
			return
		}
		walze = walze[1:] + walze[:1]
	}
	return ""
}

func inputClearWord() (word string) {
	fmt.Println("Bitte Suchwort eingeben: ")
	fmt.Scanf("%s", &word)
	return
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
	fmt.Printf("Bitte Walzenstellung eingeben (0 - %d): ", sizeWalze)
	for {
		fmt.Scanf("%d", &walzenstellung)
		if walzenstellung >= 0 || walzenstellung <= sizeWalze {
			break
		}
		fmt.Printf("Bitte Wert zwischen 0 und %d eingeben: ", sizeWalze)
	}
	for i := 0; i < walzenstellung; i++ {
		walze = walze[1:] + walze[:1]
	}
	return
}

func inputMode() (mode int) {
	fmt.Println("0 - Verschlüsseln\n1 - Entschlüsseln\n2 - Text Knacken")
	for {
		fmt.Scanf("%d", &mode)
		if mode != modeDecrypt && mode != modeEncrypt && mode != modeCrack {
			return
		}
		fmt.Println("Gültige Werte: \n0 - Verschlüsseln\n1 - Entschlüsseln\n2 - Text Knacken")
	}
}
