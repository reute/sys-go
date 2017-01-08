package main

import ("fmt"; "os";  "bufio"; "strings") 

const (
	running = iota
	noCandidatesLeft
    noCitiesLeft
    begin
    end
)

func main() {
    const filename = "cities.txt"
    cities := readFromFile(filename)
    var chain []string
    gameStatus := checkGameStatus(chain, cities)
    for gameStatus == running  {        
        fmt.Printf("%d Words left\n", len(cities))
        inputCity := strings.TrimRight(getInputCity(), "\n") 
        if inputCity == "" {            
            printHintList(findCandidates(chain, cities))
        } else {
            found, cityPos := searchCity(inputCity, cities)       
            if found {
                    fits, chainPos := checkChain(cities[cityPos], chain)
                    if fits {
                        move(&chain, &cities, cityPos, chainPos)
                        printChain(chain)                
                        gameStatus = checkGameStatus(chain, cities)                   
                    } else {
                        fmt.Println("Word doesn't fit")                         
                    }                         
            } else {
                fmt.Println("Word not found in list")
            }         
        }         
    }
    switch gameStatus {
        case noCandidatesLeft: fmt.Printf("No further candidates available, ending. %d words in Chain", len(chain)) 
        case noCitiesLeft: fmt.Printf("No further cities available, ending. %d words in Chain", len(chain)) 
    }  
}

func move(chain, cities *[]string, cityPos, chainPos int) {
    if chainPos == begin { 
        fmt.Println("begin")      
        *chain = append(*chain, "")
        copy((*chain)[1:], (*chain)[:])
        (*chain)[0] = (*cities)[cityPos] 
    } else if chainPos == end {
        fmt.Println("end")
        *chain = append(*chain, (*cities)[cityPos])
    }
    *cities = append((*cities)[:cityPos], (*cities)[cityPos+1:]...)    
}

func checkChain(city string, chain []string) (bool, int) {
    if len(chain) == 0 {
        return true, begin
    }
    if fitsBegin(city, chain[0]) {
        return true, begin
    }
    if fitsEnd(city, chain[len(chain)-1]) {
        return true, end
    }
    return false, 0
}

func fitsBegin(city, firstCityInChain string) bool { 
    city = strings.ToUpper(city)
    cityLastChar := string(city[len(city)-1])
    return strings.HasPrefix(firstCityInChain, cityLastChar)
}

func fitsEnd(city, lastCityInChain string) bool {
    lastCityInChain = strings.ToUpper(lastCityInChain) 
    city = strings.ToUpper(city) 
    cityFirstChar := string((city)[0])  
    return strings.HasSuffix(lastCityInChain, cityFirstChar)
}

func readFromFile(filename string) (cities []string) {   
    file, _ := os.Open(filename)
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines) 
    {
        var city string
        for i := 0; scanner.Scan(); i++ {
            city = scanner.Text()        
            cities = append(cities, city)
        } 
    }
    return
}

func findCandidates(chain, cities []string) ([]string) {
    if len(chain) == 0  || len(cities) == 0 { 
        return cities 
    }
    candidates := []string{}      
    firstCityInChain := chain[0]
    lastCityInChain := chain[len(chain)-1]
    for _, city := range cities {       
        if fitsBegin(city, firstCityInChain) {          
            candidates = append(candidates, city)            
        } else if fitsEnd(city, lastCityInChain) {           
            candidates = append(candidates, city)
        }   
    }     
    return candidates
}

func checkGameStatus(chain, cities []string) (int) {
    if len(cities) == 0 {
        return noCitiesLeft
    }
    if len(chain) == 0 {
        return running
    }
    candidates := findCandidates(chain, cities) 
    if len(candidates) != 0 {
        return running    
    } else {
        return noCandidatesLeft
    }
}

func getInputCity() (inputCity string) { 
    fmt.Print("Next City: ")
    tmp := bufio.NewReader(os.Stdin)
	inputCity, _ = tmp.ReadString('\n')
    inputCity = strings.TrimSpace(inputCity)
    return
}

func printHintList(candidates []string) {    
    for _, city := range candidates {
            fmt.Printf("%s, ", city) 
    }
    fmt.Println()
}

func searchCity(inputCity string, cities []string) (bool, int) {
    for i, city := range cities {
		if city == inputCity {
            return true, i
        }
	}
    return false, 0  
}

func printChain(chain []string) {
    if len(chain) != 0 {
        fmt.Print(chain[0])
        fmt.Print(" ... ")
        fmt.Print(chain[len(chain)-1])
        fmt.Println()
    }
}