package main

import (
	"bufio"
	"fmt"
	"os"
)

// 输出有重复行的文件名称
func main() {
	fileName := []string{}
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, fileName)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, fileName)
			f.Close()

		}
	}
	for _, v := range fileName {
		fmt.Println(v)
	}
}
func countLines(f *os.File, fileName []string) []string {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	for _, n := range counts {
		if n > 1 {
			fileName = append(fileName, f.Name())
		}
	}
	return fileName
}
