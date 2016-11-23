package main

import (
	"fmt"
)

type spieler struct {
    zahl, einsatz, gewinn, kontostand int
}

func (s *spieler) addWins() {    
    s.kontostand += s.gewinn 
}

func (s *spieler) win() {
    s.gewinn += s.einsatz
}

func (s *spieler) lose() {
    s.kontostand -= s.einsatz
}

func (s *spieler) inputEinsatz() {  
    for {
        fmt.Print("Ihr Einsatz: ")    
        fmt.Scanf("%d", &s.einsatz) 
        if s.einsatz > s.kontostand {
            fmt.Printf("Sie haben nicht genügend Geld !!\n")
            continue
        } else if s.einsatz < 0 {
            fmt.Printf("Einsatz muss über 0 sein !!\n")	
            continue
        } else {
            break 
        }       
    }   
}

func (s *spieler) inputZahl() {
    for {
        fmt.Printf("Ihre Zahl: ")
        fmt.Scanf("%d", &s.zahl)
        if s.zahl > 6 {
            fmt.Printf("Das ist zu hoch !!\n")	
            continue	
        } else if s.zahl < 1 {
            fmt.Printf("Das ist zu niedrig !!\n")
            continue
        } else {
            break
        }
	}
}