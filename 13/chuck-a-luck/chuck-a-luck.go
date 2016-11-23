package main

import (
	"fmt"
	"math/rand"
)

var s = spieler{kontostand: 1000}

func main() { 
    rand.Seed(42)
    fmt.Println("**** Chuck-a-luck ****\nIn jeder Runde können Sie einen Teil davon auf eine der Zahlen 1 bis 6 setzen. Dann werden 3 Würfel geworfen. Falls Ihr Wert dabei ist, erhalten Sie Ihren Einsatz zurück und zusätzlich Ihren Einsatz für jeden Würfel, der die von Ihnen gesetzte Zahl aufweist")
    gameloop()
}

func gameloop() {
    for {
        fmt.Printf("Sie haben %d Geldeinheiten\n", s.kontostand)
        s.inputEinsatz()
        if s.einsatz == 0 {
            fmt.Printf("Glückwunsch, Sie verlassen das Casino mit %d Geldeinheiten !!\n", s.kontostand)
            break 
        }        
        s.inputZahl() 
        rollDice(&s)

		if s.gewinn == 0 {
            s.lose()
			fmt.Printf("Pech, da war nichts fuer Sie dabei!!\n")
		} else {
            s.addWins()	
			fmt.Printf("Glückwunsch, Sie erhalten %d Geldeinheiten!!\n", s.gewinn)
		}

		if s.kontostand <= 0 {
			fmt.Printf("Du bist leider Pleite !!!\n")
			break
		}
	}
}

func rollDice(s *spieler) {
    var tmp int
    s.gewinn = 0
    fmt.Printf("Die Würfel sind gefallen: ")
    for i := 1; i <= 3; i++ {
        tmp = rand.Intn(6)
        fmt.Printf("%d ",tmp)
        if tmp == s.zahl {
            s.win()
        }
    }    
}