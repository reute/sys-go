package main

import (
  "os"
  "fmt"
  "bufio"
   "io/ioutil"
   "regexp"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    stdin, _ := ioutil.ReadAll(reader)
    text := string(stdin)
    // fmt.Println("Something on STDIN: " + text)
 	re := regexp.MustCompile(`http://[^"]*`)    
	result := re.FindAllString(text, -1)
    for i, val := range(result) {
      fmt.Printf(" %d: %s\n",i, val)
    }

}