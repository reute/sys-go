package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	clear        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "
	defaultWalze = "EKLMF6GDQVZ0TO8Y XUSP2IB4CJ5AR197W3NH"
	sizeWalze    = len(clear)
)

const (
	modeEncrypt = iota + 1
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
	if mode == modeCrack {
		word := inputClearWord()
		result := crack(text, word)
		if result == "" {
			fmt.Println("Suchwort nicht gefunden !")
		} else {
			fmt.Printf("Gecrackter Text:\n%s\n", result)
		}
		return
	}
	// decrypt or encrypt
	walze := getWalze()
	e := enigma{text, walze}
	if mode == modeDecrypt {
		fmt.Printf("\nEntschlüsselter Text:\n%s\n", e.decrypt())
	} else {
		fmt.Printf("\nVerschlüsselter Text:\n%s\n", e.encrypt())
	}
}

func inputText() string {
	fmt.Println("Bitte Text eingeben:")
	in := bufio.NewReader(os.Stdin)
	text, _ := in.ReadString('\n')
	text = strings.TrimSpace(text)
	text = strings.ToUpper(text)
	return text
}

func inputMode() (mode uint) {
	text := fmt.Sprintf("%d - Verschlüsseln\n%d - Entschlüsseln\n%d - Text Knacken\n", modeEncrypt, modeDecrypt, modeCrack)
	fmt.Print(text)
	for {
		fmt.Scanf("%d", &mode)
		switch mode {
		case modeEncrypt, modeDecrypt, modeCrack:
			return
		default:
			fmt.Print("Gültige Werte: " + text)
		}
	}
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
	walze = walze[walzenstellung:] + walze[:walzenstellung]
	return
}

func (e *enigma) encrypt() (result string) {
	for _, runechar := range e.text {
		if !strings.ContainsRune(clear, runechar) {
			runechar = ' '
		}
		i := strings.IndexRune(clear, runechar)
		result += string(e.walze[i])
		e.walze = e.walze[1:] + e.walze[:1]
	}
	return
}

func (e *enigma) decrypt() (result string) {
	for _, runechar := range e.text {
		i := strings.IndexRune(e.walze, runechar)
		result += string(clear[i])
		e.walze = e.walze[1:] + e.walze[:1]
	}
	return
}

func crack(encrypted string, knownWord string) (decrypted string) {
	var e enigma
	walze := defaultWalze
	for walzenstellung := 0; walzenstellung < sizeWalze; walzenstellung++ {
		e = enigma{encrypted, walze}
		decrypted = e.decrypt()
		if strings.Contains(decrypted, knownWord) {
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
