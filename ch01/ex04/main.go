package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countLines(f *os.File, counts map[string]int, fileNames map[string][]string) {
	input := bufio.NewScanner(f)
	fileName := f.Name()
	for input.Scan() {
		counts[input.Text()]++
		if !contains(fileNames[input.Text()], fileName) {
			fileNames[input.Text()] = append(fileNames[input.Text()], fileName)
		}
	}
}

func contains(slice []string, target string) bool {
	for _, val := range slice {
		if val == target {
			return true
		}
	}
	return false
}

func main() {
	fileNames := make(map[string][]string)
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(fileNames[line], " "))
		}
	}
}
