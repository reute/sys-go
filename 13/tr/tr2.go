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
    fmt.Println(content)  
    fmt.Println(chars1)    
    fmt.Println(chars2)  
    newCharset := createNewCharset(chars1, chars2)
    fmt.Println("charset", charset)
    fmt.Println("newcharset:", newCharset)
    newContent := replaceChars(content, newCharset)
    fmt.Println("newContent", newContent)
    writeFile(filename, newContent)
}

func createNewCharset(chars1, chars2 string) (newCharset string) {
    tmp := []byte(charset)
    icharset := 0
    for ichars2, val := range chars1 {
        icharset = strings.IndexRune(charset, val)
        tmp[icharset] = chars2[ichars2]
    }
    newCharset = string(tmp)
    return 
}

func replaceChars(content, newCharset string) (newContent []byte) {    
    var icharset int
    newContent = []byte(content)
    for icon, charcon := range content {
        icharset = strings.IndexRune(charset, charcon)
      	if  icharset != -1 {
            newContent[icon] = newCharset[icharset]
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