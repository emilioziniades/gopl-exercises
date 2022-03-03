package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	fileWithDups := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileWithDups)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileWithDups)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			dupFiles := fileWithDups[line]
			files := make([]string, 0)
			for k := range dupFiles {
				files = append(files, k)

			}
			fmt.Printf("%d\t%s\t(%v)\n", n, line, strings.Join(files, ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, dups map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if _, exists := dups[input.Text()]; !exists {
			dups[input.Text()] = make(map[string]bool)
		}
		dups[input.Text()][f.Name()] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}
