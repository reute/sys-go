package main

import ("os";"fmt"; "container/list"; "io/ioutil")

const MAX_FILES = 10

type file struct {
    size int64
    name string
    path string
}

var fileList *list.List

func main() {
    fileList = list.New()
    addInitialFile()
    if len(os.Args) > 1 {
        scanDir(os.Args[1])
    } else {
        scanDir(".")
    }
    printList()
}

func addFile(newFile file) {
    var inserted bool
    for e := fileList.Front(); e != nil; e = e.Next() {
        elSize := ((e.Value).(file)).size
        if elSize >= newFile.size {
            fileList.InsertBefore(newFile, e)
            inserted = true
            break
        }
    } 
    if !inserted {
        fileList.PushBack(newFile)
    }
    if fileList.Len() == MAX_FILES + 1 {
       fileList.Remove(fileList.Front()) 
    } 
}

func printList() { 
    fmt.Println("Result: ")
    for e := fileList.Back(); e != nil; e = e.Prev() {                   
        fmt.Println(e.Value)
    }
}

func addInitialFile() {
    addFile(file{0, "dummy", "."})
}

func scanDir(path string) {
    items, _ := ioutil.ReadDir(path)
    for _, item := range items {
        itemName := item.Name()
        if item.IsDir() {
            scanDir(path + "/" + itemName)
        } else  {
            itemSize := item.Size()
            elSize := ((fileList.Front().Value).(file)).size 
            if itemSize > elSize {
                addFile(file{itemSize, item.Name(), path})
            }           
        }
    } 
}
