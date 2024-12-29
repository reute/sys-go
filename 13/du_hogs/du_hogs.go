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

func getDirName() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return "."
}

func scanDir(path string) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory: %v\n", err)
		return
	}
	for _, dirEntry := range dirEntries {
		name := dirEntry.Name()
		if dirEntry.IsDir() {
			scanDir(filepath.Join(path, name))
		} else {
			info, err := dirEntry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error getting file info: %v\n", err)
				continue
			}
			size := info.Size()
			addFile(file{size, name, path})
		}
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

func printList() {
	fmt.Println("Result:")
	for _, f := range fileList {
		fmt.Printf("%s (%d kbytes)\n", filepath.Join(f.path, f.name), f.size/1024)
	}
}
