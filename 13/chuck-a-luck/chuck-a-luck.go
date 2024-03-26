package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type spieler struct {
	kontostand int
}

func (s *spieler) win(gewinn int) {
	s.kontostand += gewinn
}

func (s *spieler) lose(einsatz int) {
	s.kontostand -= einsatz
}

const (
	beenden = iota
	laufend
	pleite
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	kontostand := 1000
	if len(os.Args) > 1 {
		arg, _ := strconv.Atoi(os.Args[1])
		if arg > 0 {
			kontostand = arg
		}
	}
	s1 := spieler{kontostand: kontostand}
	var zahl, gewinn, einsatz int
	spielStatus := laufend
	fmt.Println("**** Chuck-a-luck ****\nIn jeder Runde können Sie einen Teil davon auf eine der Zahlen 1 bis 6 setzen. Dann werden 3 Würfel geworfen. Falls Ihr Wert dabei ist, erhalten Sie Ihren Einsatz zurück und zusätzlich Ihren Einsatz für jeden Würfel, der die von Ihnen gesetzte Zahl aufweist")

	for spielStatus == laufend {
		fmt.Printf("Sie haben %d Geldeinheiten\n", s1.kontostand)
		einsatz = inputEinsatz(s1.kontostand)
		if einsatz == 0 {
			spielStatus = beenden
			continue
		}
		zahl = inputZahl()
		gewinn = berechneGewinn(einsatz, zahl)
		if gewinn == 0 {
			s1.lose(einsatz)
			fmt.Printf("Pech, da war nichts fuer Sie dabei!!\n")
		} else {
			s1.win(gewinn)
			fmt.Printf("Glückwunsch, Sie erhalten %d Geldeinheiten!!\n", gewinn)
		}
		if s1.kontostand <= 0 {
			spielStatus = pleite
			continue
		}
	}
	switch spielStatus {
	case beenden:
		fmt.Printf("Glückwunsch, Sie verlassen das Casino mit %d Geldeinheiten !!\n", s1.kontostand)
	case pleite:
		fmt.Printf("Sie sind leider Pleite !!!\n")
	}
}

func berechneGewinn(einsatz, zahl int) (gewinn int) {
	var tmp int
	fmt.Printf("Die Würfel sind gefallen: ")
	for i := 1; i <= 3; i++ {
		tmp = rand.Intn(6)
		fmt.Printf("%d ", tmp)
		if tmp == zahl {
			gewinn += einsatz
		}
	}
	return
}

func inputZahl() (zahl int) {
	for {
		fmt.Printf("Ihr Tip: ")
		fmt.Scanf("%d", &zahl)
		if zahl > 6 {
			fmt.Printf("Das ist zu hoch !!\n")
			continue
		} else if zahl < 1 {
			fmt.Printf("Das ist zu niedrig !!\n")
			continue
		} else {
			return
		}
	}
}

func inputEinsatz(kontostand int) (einsatz int) {
	for {
		fmt.Print("Ihr Einsatz: ")
		fmt.Scanf("%d", &einsatz)
		if einsatz > kontostand {
			fmt.Printf("Sie haben nicht genügend Geld !!\n")
			continue
		} else if einsatz < 0 {
			fmt.Printf("Einsatz muss über 0 sein !!\n")
			continue
		} else {
			break
		}
	}
	return
}
