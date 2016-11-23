package main

import ("fmt"; "os"; "bufio"; "math/rand"; "strings"; "strconv") 	

const MTN_MAX = 8
const FILENAME = "berge"

type mountain struct {
    name string
    height int
}

func main() {
    var from, to int       
    rand.Seed(42) 
    unsorted := readFile()  
    sorted := make([]mountain, 0)
    validMove := true   
    // gameloop
    for validMove && len(unsorted) > 0  {
        fmt.Println("Current state:")
        printMountains(unsorted)
        fmt.Println("Still to be sorted:")
        printMountains(sorted)
        from, to = inputUser()
        unsorted, sorted = move(from, unsorted, to, sorted)  
        validMove = isSorted(sorted)
    }
    
    if validMove {
       fmt.Println("You have  won !!")  
    } else {
       fmt.Println("Sorry, then it is no longer sorted:")       
    }
    
    for _, mtn := range(sorted) {
        fmt.Printf(" %5d: %s \n", mtn.height, mtn.name)
    }
    fmt.Println("Bye!")
    fmt.Printf("You got %d Points.\n", len(sorted) - 1)  
}

func move(from int, unsorted []mountain, to int, sorted []mountain) ([]mountain, []mountain) {
    sorted = append(sorted, mountain{})
    copy(sorted[to+1:], sorted[to:])
    sorted[to] = unsorted[from]
    unsorted = append(unsorted[:from], unsorted[from+1:]...)  
    return unsorted, sorted  
}

func inputUser() (int, int) {
    fmt.Println("What is to be inserted where? ")
    in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
    strSlice := strings.Fields(str)
    from, _ := strconv.Atoi(strSlice[0])
    to, _ := strconv.Atoi(strSlice[1])  
    return from, to  
}

func isSorted(mountains []mountain) bool {
    var tmp, hightest int
    for _, mtn := range(mountains) {
        tmp = mtn.height
        if tmp > hightest {
            hightest = tmp
        } else {
            return false
        }
    }
    return true
}

func readFile() []mountain { 
    var unsorted [MTN_MAX]mountain
    inFile, _ := os.Open(FILENAME)
    defer inFile.Close()
    scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines) 
    var tmp []string
    var name string
    var height, r int
    for i := 0; scanner.Scan(); i++ {
        tmp = strings.Split(scanner.Text(), ":")
        name = tmp[0]
        height, _ = strconv.Atoi(tmp[1])         
        if i < MTN_MAX {
           unsorted[i] = mountain{name, height}
        } else {
            r = rand.Intn(i)
            if r < MTN_MAX {
                unsorted[r] = mountain{name, height}
            }
        }
    }
    return unsorted[:]
}

func printMountains(mountains []mountain) {
    for i, mtn := range(mountains) {
        fmt.Printf(" %d: %s\n", i, mtn.name)
    }
}