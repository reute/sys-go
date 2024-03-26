package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const spielfeldbreite = 23
const (
	spieler = iota
	computer
	laufend
	spielerGewinnt
	computerGewinnt
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("*** Drei Kreuze ***\nGegeben ist eine Kette von 23 freien Feldern. In jedem Zug setzt jeder der Spieler ein X auf ein freies Feld. Wenn dadurch drei oder mehr X benachbart sind, hat der Spieler gewonnen.")
	var amZug, spielfeld uint
	zeichneSpielfeld(spielfeld)
	var anfangen string
	fmt.Print("Wollen Sie anfangen? (j/n):  ")
	fmt.Scanf("%s", &anfangen)
	if anfangen == "j" {
		amZug = spieler
	} else {
		amZug = computer
	}
	var spielstatus uint = laufend
	for spielstatus == laufend {
		neuerZug(amZug, &spielfeld)
		zeichneSpielfeld(spielfeld)
		spielstatus = pruefeSpielstand(spielfeld, amZug)
		amZug = (amZug + 1) % 2
	}
	if spielstatus == spielerGewinnt {
		fmt.Println("Der Spieler gewinnt !")
	} else {
		fmt.Println("Der Computer gewinnt !")
	}
}

func neuerZug(amZug uint, spielfeld *uint) {
	var zug uint
	if amZug == computer {
		for {
			zug = uint(rand.Intn(int(spielfeldbreite)))
			if feldBelegen(zug, spielfeld) {
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
			if feldBelegen(zug, spielfeld) {
				break
			} else {
				fmt.Println("Feld ist schon belegt !")
			}
		}
	}
}

func zeichneSpielfeld(spielfeld uint) {
	var i uint
	for i = 0; i < spielfeldbreite; i++ {
		if feldBelegt(i, spielfeld) {
			fmt.Print("\\/ ")
		} else {
			fmt.Print("   ")
		}
	}
	fmt.Println()
	for i = 0; i < spielfeldbreite; i++ {
		if feldBelegt(i, spielfeld) {
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

func feldBelegen(pos uint, spielfeld *uint) bool {
	if feldBelegt(pos, *spielfeld) {
		return false
	}
	var bit uint = 1
	bit = bit << pos
	*spielfeld = *spielfeld | bit
	return true
}

func feldBelegt(pos uint, spielfeld uint) bool {
	var bit uint = 1
	bit = bit << pos
	if (bit & spielfeld) == bit {
		return true
	}
	return false
}

func pruefeSpielstand(spielfeld, amZug uint) (spielstatus uint) {
	var kreuze, i uint
	spielstatus = laufend
	for i = 0; i < spielfeldbreite; i++ {
		if feldBelegt(i, spielfeld) {
			kreuze++
		} else {
			kreuze = 0
		}
		if kreuze == 3 {
			if amZug == spieler {
				spielstatus = spielerGewinnt
			} else {
				spielstatus = computerGewinnt
			}
			break
		}
	}
	return
}
