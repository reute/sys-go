package main

import (
    "os"
    "io/ioutil"
	"fmt"
    "strings"
)

const (
    charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func main() {    
    var filename, chars1, chars2 string
    if len(os.Args) > 3 {        
        chars1 = os.Args[1]
        chars2 = os.Args[2]
        filename = os.Args[3]
    }
    content := readFile(filename) 
    fmt.Println(chars1)    
    fmt.Println(chars2) 
    fmt.Println(content)  
    chars1 = expandRange(chars1)
    chars2 = expandRange(chars2)
    fmt.Println("chars1", chars1)
    fmt.Println("chars2", chars2)
    newCharset := createNewCharset(chars1, chars2)
    fmt.Println("charset", charset)
    fmt.Println("newcharset:", newCharset)
    newContent := replaceChars(content, newCharset)
    writeFile(filename, newContent)
}

func expandRange(chars string) string {
    expChars := make([]byte, 0, 256)
    var from, to byte
    for i := 0; i < len(chars); i++ {        
        if chars[i] == '-' {
            from = chars[i-1]
            to = chars[i+1]
            // fmt.Println("from", from)
            // fmt.Println("to", to)
            from++       
            for i := from; i <= to; i++ {
                expChars = append(expChars, i)                
            }
            i++
        } else {
            expChars = append(expChars, chars[i])
        }        
    }
    return string(expChars)
}

func createNewCharset(chars1, chars2 string) (newCharset string) {
    tmp := []byte(charset)
    iet := 0
    for i2, val := range chars1 {
        iet = strings.IndexRune(charset, val)
        tmp[iet] = chars2[i2]
    }
    newCharset = string(tmp)
    return 
}

func replaceChars(content, newCharset string) (newContent []byte) {    
    var iet int
    newContent = []byte(content)
    for icon, charcon := range content {
        iet = strings.IndexRune(charset, charcon)
      	if  iet != -1 {
            newContent[icon] = newCharset[iet]
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