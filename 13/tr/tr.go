package main

import (
    "os"
    "io/ioutil"
	"fmt"
    "strings"
    "math"
)

func main() {    
    var filename, chars1, chars2 string
    if len(os.Args) > 3 {        
        chars1 = os.Args[1]
        chars2 = os.Args[2]
        filename = os.Args[3]
    }
    text := readFile(filename) 
    fmt.Println(chars1)    
    fmt.Println(chars2) 
    fmt.Println(text)  
    chars1 = expandRange(chars1)
    chars2 = expandRange(chars2)
    fmt.Println("chars1", chars1)
    fmt.Println("chars2", chars2)
    newCharset := createNewCharset([]byte(chars1), []byte(chars2))
    fmt.Println("newcharset:", newCharset)
    newText := replaceChars(text, newCharset)
    writeFile(filename, newText)
}

func expandRange(chars string) string {
    expChars := make([]byte, 0, math.MaxInt8)
    var from, to byte
    for i := 0; i < len(chars); i++ {        
        if chars[i] == '-' {
            from = chars[i-1]
            to = chars[i+1]
            // fmt.Println("from", from)
            // fmt.Println("to", to)               
            for i := from+1; i <= to; i++ {
                expChars = append(expChars, i)                
            }
            i++
        } else {
            expChars = append(expChars, chars[i])
        }        
    }
    return string(expChars)
}

func createNewCharset(chars1, chars2 []byte) (newCharset [math.MaxInt8]byte) {
    for indexChars1, val := range chars1 {       
        newCharset[val] = chars2[indexChars1]
    }
    return 
}

func replaceChars(text string, newCharset [math.MaxInt8]byte) (newText []byte) {
    newText = []byte(text)
    for indexContent, val := range text {       
      	if  newCharset[val] != 0 {
            newText[indexContent] = newCharset[val]
        }	
    }
    return 
}

func readFile(filename string) string {
    tmp, _ := ioutil.ReadFile(filename)    
    content := strings.TrimRight(string(tmp), "\n") 
    return string(content)
}

func writeFile(filename string, content []byte) error {
    return ioutil.WriteFile(filename, content, 0644)
}