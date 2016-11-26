package main

import (
	"fmt"
	"math/rand"
    "time"
)

type spieler struct {
    zahl, einsatz, kontostand int
}

func (s *spieler) win(gewinn int) {    
    s.kontostand += gewinn 
}

func (s *spieler) lose() {
    s.kontostand -= s.einsatz
}

var s = spieler{kontostand: 1000}

const (
	beenden = iota
	pleite
)

func main() { 
    rand.Seed(time.Now().UTC().UnixNano())
    fmt.Println("**** Chuck-a-luck ****\nIn jeder Runde können Sie einen Teil davon auf eine der Zahlen 1 bis 6 setzen. Dann werden 3 Würfel geworfen. Falls Ihr Wert dabei ist, erhalten Sie Ihren Einsatz zurück und zusätzlich Ihren Einsatz für jeden Würfel, der die von Ihnen gesetzte Zahl aufweist")
    result := gameloop()
    exit(result)
}

func exit(result int) {
    if result == beenden {
         fmt.Printf("Glückwunsch, Sie verlassen das Casino mit %d Geldeinheiten !!\n", s.kontostand)
    } else {
        fmt.Printf("Du bist leider Pleite !!!\n")
    }

}

func gameloop() int {
    var gewinn int
    for {
        fmt.Printf("Sie haben %d Geldeinheiten\n", s.kontostand)
        s.einsatz = inputEinsatz()
        if s.einsatz == 0 {       
            return beenden
        }        
        s.zahl = inputZahl() 
        gewinn = rollDice(&s)
		if gewinn == 0 {
            s.lose()
			fmt.Printf("Pech, da war nichts fuer Sie dabei!!\n")
		} else {
            s.win(gewinn)	
			fmt.Printf("Glückwunsch, Sie erhalten %d Geldeinheiten!!\n", gewinn)
		}

		if s.kontostand <= 0 {	
			return pleite
		}
	}
}

func rollDice(s *spieler) int {
    var tmp, gewinn int
    fmt.Printf("Die Würfel sind gefallen: ")
    for i := 1; i <= 3; i++ {
        tmp = rand.Intn(6)
        fmt.Printf("%d ",tmp)
        if tmp == s.zahl {
            gewinn += s.einsatz
        }
    } 
    return gewinn   
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

func inputEinsatz() (einsatz int) {  
    for {
        fmt.Print("Ihr Einsatz: ")    
        fmt.Scanf("%d", &einsatz) 
        if einsatz > s.kontostand {
            fmt.Printf("Sie haben nicht genügend Geld !!\n")
            continue
        } else if einsatz < 0 {
            fmt.Printf("Einsatz muss über 0 sein !!\n")	
            continue
        } else {
            return 
        }       
    }   
}