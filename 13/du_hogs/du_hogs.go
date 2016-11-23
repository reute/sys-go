package main

import ("os";"fmt"; "container/list"; "io/ioutil")

const MAX_FILES = 10

type file struct {
    size int64
    name string
    path string
}

var file_list *list.List

func main() {
    file_list = list.New()
    add_initial_file()
    if len(os.Args) > 1 {
        scan_dir(os.Args[1])
    } else {
        scan_dir(".")
    }
    printList()
}

func add_file(new_file file) {
    var inserted bool
    for e := file_list.Front(); e != nil; e = e.Next() {
        e_size := ((e.Value).(file)).size
        if e_size >= new_file.size {
            file_list.InsertBefore(new_file, e)
            inserted = true
            break
        }
    } 
    if !inserted {
        file_list.PushBack(new_file)
    }
    if file_list.Len() == MAX_FILES + 1 {
       file_list.Remove(file_list.Front()) 
    } 
}

func printList() { 
    fmt.Println("Result: ")
    for e := file_list.Back(); e != nil; e = e.Prev() {                   
        fmt.Println(e.Value)
    }
}

func add_initial_file() {
    add_file(file{0, "dummy", "."})
}

func scan_dir(path string) {
    items, _ := ioutil.ReadDir(path)
    for _, item := range items {
        item_name := item.Name()
        if item.IsDir() {
            scan_dir(path + "/" + item_name)
        } else  {
            item_size := item.Size()
            e_size := ((file_list.Front().Value).(file)).size 
            if item_size > e_size {
                add_file(file{item_size, item.Name(), path})
            }           
        }
    } 
}
