package main

import (
	"fmt"
	"os"
	"math/rand"
)

var spielfeld, letzterZug uint

const WORD_BIT = 23
const (
	Spieler = iota
	Computer
)

func zeichneSpielfeld() {
	var i uint
	for i = 0; i < WORD_BIT; i++ {
		if feldBelegt(i) {
			fmt.Print("\\/ ")
		} else {
			fmt.Print("   ")
		}		
	}
	fmt.Println()

	for i = 0; i < WORD_BIT; i++ {
		if feldBelegt(i) {
			fmt.Print("/\\ ")
		} else {
			fmt.Print("   ")
		}	
	}
	fmt.Println()

	for i = 0; i < WORD_BIT; i++ {
		fmt.Printf("%02d ", i)
	}
	fmt.Println()
}

func feldBelegen(pos uint) (bool) {
	if feldBelegt(pos) {
		return false
	}
	var bit uint = 1
	bit = bit << pos
	spielfeld = spielfeld | bit
	return true
}

func feldBelegt(pos uint) (bool) {
	var bit uint = 1
	bit = bit << pos
	if (bit & spielfeld) == bit {
		return true
	} else {
		return false
	}
}

func zugSpieler() {
	letzterZug = Spieler
	var zug uint	
	for {
		fmt.Print("Ihr Zug : ")
		fmt.Scanf("%d", &zug)
		if zug == 99 {
			fmt.Println("Program wird beendet !")
			os.Exit(0)
		}

		if zug < 0 || zug > WORD_BIT {
			 fmt.Printf("Bitte Zahl zwischen 0 und %d eingeben\n", WORD_BIT-1)
			 continue
		}	
		if feldBelegen(zug) {
            break
        } else {
            fmt.Println("Feld ist schon belegt !")  
        }
	}	
}

func zugComputer() {
	letzterZug = Computer
	var pos uint	
	for {
		pos  = uint(rand.Intn(int(WORD_BIT)))
		if feldBelegen(pos) {
			break
		}
	}
	fmt.Printf("Der Computer belegt das Feld %d\n", pos)
}

func pruefeSpielstand() {
	var kreuze uint;
	var i uint;
	for i = 0; i < WORD_BIT; i++ {
    	if feldBelegt(i) {
			kreuze++
    	} else {
        	kreuze = 0    
   		} 
   		if kreuze == 3 {
   			if letzterZug == Spieler {
   				fmt.Println("Der Spieler gewinnt !")
   			} else {
				fmt.Println("Der Computer gewinnt !")
   			}
   			os.Exit(0)
   		}   	
	}
}

func main() {
	var anfangen string
	rand.Seed(42)

	fmt.Println("*** Drei Kreuze ***\nGegeben ist eine Kette von 23 freien Feldern. In jedem Zug setzt jeder der Spieler ein X auf ein freies Feld. Wenn dadurch drei oder mehr X benachbart sind, hat der Spieler gewonnen.")
	zeichneSpielfeld()
	fmt.Print("Wollen Sie anfangen? (j/n):  ")
	fmt.Scanf("%s", &anfangen)

	if anfangen == "j" {
        zugSpieler()
        zeichneSpielfeld()
        pruefeSpielstand()
	} 
	for {
        zugComputer()
        zeichneSpielfeld()
        pruefeSpielstand()
        zugSpieler()
        zeichneSpielfeld()
        pruefeSpielstand()
	}
}