package main

import (
	"fmt"
	"os"
	"math/rand"
    "time"
)

const spielfeldbreite = 23
const (
	spieler = iota
	computer
)
var amZug = computer
var spielfeld uint

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("*** Drei Kreuze ***\nGegeben ist eine Kette von 23 freien Feldern. In jedem Zug setzt jeder der Spieler ein X auf ein freies Feld. Wenn dadurch drei oder mehr X benachbart sind, hat der Spieler gewonnen.")
	zeichneSpielfeld()
	var anfangen string
	fmt.Print("Wollen Sie anfangen? (j/n):  ")
	fmt.Scanf("%s", &anfangen)
	if anfangen == "j" {
		amZug = spieler
	} 
	gameloop()
	exit()
}

func gameloop() {
	for {
		neuerZug()
        zeichneSpielfeld()
		if pruefeSpielstand() != true {
			break
		}	 
		amZug = (amZug + 1) % 2
	}
}

func exit() {
	if amZug == spieler {
		fmt.Println("Der Spieler gewinnt !")
	} else {
		fmt.Println("Der Computer gewinnt !")
	}
}

func neuerZug() {
	var zug uint
	if amZug == computer {		
		for {
			zug  = uint(rand.Intn(int(spielfeldbreite)))
			if feldBelegen(zug) {
				break
			}
		}
		fmt.Printf("Der Computer belegt das Feld %d\n", zug)
	} else {		
		for {
			fmt.Print("Ihr Zug : ")
			fmt.Scanf("%d", &zug)
			if zug == 99 {
				fmt.Println("Program wird beendet !")
				os.Exit(0)
			}
			if zug < 0 || zug > spielfeldbreite {
				fmt.Printf("Bitte Zahl zwischen 0 und %d eingeben\n", spielfeldbreite-1)
				continue
			}	
			if feldBelegen(zug) {
				break
			} else {
				fmt.Println("Feld ist schon belegt !")  
			}
		}
	}
}

func zeichneSpielfeld() {
	var i uint
	for i = 0; i < spielfeldbreite; i++ {
		if feldBelegt(i) {
			fmt.Print("\\/ ")
		} else {
			fmt.Print("   ")
		}		
	}
	fmt.Println()

	for i = 0; i < spielfeldbreite; i++ {
		if feldBelegt(i) {
			fmt.Print("/\\ ")
		} else {
			fmt.Print("   ")
		}	
	}
	fmt.Println()

	for i = 0; i < spielfeldbreite; i++ {
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

func pruefeSpielstand() bool {
	var kreuze uint;
	var i uint;
	for i = 0; i < spielfeldbreite; i++ {
    	if feldBelegt(i) {
			kreuze++
    	} else {
        	kreuze = 0    
   		} 
   		if kreuze == 3 {
			return false   			
   		}   	
	}
	return true
}