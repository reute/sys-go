package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

const MaxCharsetSize = math.MaxInt8

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: <chars1> <chars2> <filename>")
		return
	}
	search := os.Args[1]
	replace := os.Args[2]
	filename := os.Args[3]

	search = expandRange(search)
	replace = expandRange(replace)
	if len(search) != len(replace) {
		fmt.Println("Error: search and replace chars must have the same length")
		return
	}

	newCharset := createNewCharset([]byte(search), []byte(replace))
	err := processFile(filename, newCharset)
	if err != nil {
		fmt.Println("Error processing file:", err)
		return
	}
}

func expandRange(chars string) string {
	var builder strings.Builder
	for i := 0; i < len(chars); i++ {
		if chars[i] == '-' && i > 0 && i < len(chars)-1 {
			from := chars[i-1]
			to := chars[i+1]
			for j := from + 1; j <= to; j++ {
				builder.WriteByte(j)
			}
			i++
		} else {
			builder.WriteByte(chars[i])
		}
	}
	return builder.String()
}

func createNewCharset(search, replace []byte) (newCharset [MaxCharsetSize]byte) {
	for index, codepoint := range search {
		newCharset[codepoint] = replace[index]
	}
	return
}

func processFile(filename string, newCharset [MaxCharsetSize]byte) error {
	input, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.CreateTemp("", "tr-cp-*")
	if err != nil {
		return err
	}
	defer output.Close()

	reader := bufio.NewReader(input)
	writer := bufio.NewWriter(output)

	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if newChar := newCharset[char]; newChar != 0 {
			_, err = writer.WriteRune(rune(newChar))
		} else {
			_, err = writer.WriteRune(char)
		}
		if err != nil {
			return err
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	if err := os.Rename(output.Name(), filename); err != nil {
		return copyFileContents(output.Name(), filename)
	}

	return nil
}

// Helper function to copy file contents
func copyFileContents(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Sync()
}
