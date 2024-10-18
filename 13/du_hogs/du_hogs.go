package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

const MAX_FILES = 10

type file struct {
	size int64
	name string
	path string
}

var fileList []file

func main() {
	scanDir(getDirName())
	printList()
}

func scanDir(path string) {
	items, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory: %v\n", err)
		return
	}
	for _, item := range items {
		itemName := item.Name()
		if item.IsDir() {
			scanDir(filepath.Join(path, itemName))
		} else {
			itemInfo, err := item.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error getting file info: %v\n", err)
				continue
			}
			addFile(file{itemInfo.Size(), itemName, path})
		}
	}
}

func getDirName() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return "."
}

func printList() {
	fmt.Println("Result:")
	for _, f := range fileList {
		fmt.Printf("%s (%d kbytes)\n", filepath.Join(f.path, f.name), f.size/1024)
	}
}

func addFile(newFile file) {
	fileList = append(fileList, newFile)
	sort.Slice(fileList, func(i, j int) bool {
		return fileList[i].size > fileList[j].size
	})
	if len(fileList) > MAX_FILES {
		fileList = fileList[:MAX_FILES]
	}
}
