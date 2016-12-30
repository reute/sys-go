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
    taken := make([]*mountain, 0, MTN_MAX)
    left := make([]*mountain, MTN_MAX)
    arrayMountains := readFile()
    for i := 0; i < len(arrayMountains); i++ {
        left[i] = &arrayMountains[i]
    }  
    // gameloop
    for isSorted(taken) && len(left) != 0  {
        fmt.Println("Current state:")
        printMountains(taken)
        fmt.Println("Still to be sorted:")
        printMountains(left)
        from, to = inputUser()
        taken, left = move(from, left, to, taken)   
    }
    if len(left) == 0 {        
       fmt.Println("You have  won !!")  
    } else {
       fmt.Println("Sorry, then it is no longer sorted:")       
    }
    for _, element := range(taken) {
        fmt.Printf(" %5d: %s \n", (*element).height, (*element).name)
    }
    fmt.Println("Bye!")
    fmt.Printf("You got %d Points.\n", len(taken) - 1) 
}

func isSorted(mountains []*mountain) bool {
    var tmp, hightest int
    for _, element := range(mountains) {
        tmp = (*element).height
        if tmp > hightest {
            hightest = tmp
        } else {
            return false
        }
    }
    return true
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

func readFile() (arrayMountains [MTN_MAX]mountain) {   
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
           arrayMountains[i] = mountain{name, height}
        } else {
            r = rand.Intn(i)
            if r < MTN_MAX {
                arrayMountains[r] = mountain{name, height}
            }
        }
    }
    return
}

func move(from int, left []*mountain, to int, taken []*mountain) ([]*mountain, []*mountain) {
    taken = append(taken, &mountain{})
    copy(taken[to+1:], taken[to:])
    taken[to] = left[from]
    left = append(left[:from], left[from+1:]...)    
    return taken, left
}

func printMountains(mountains []*mountain) {
    for index, element := range(mountains) {
        fmt.Printf(" %d: %s\n", index, (*element).name)
    }
}