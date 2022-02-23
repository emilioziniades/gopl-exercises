package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	wordcounts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}

	for input.Scan() {
		word := input.Text()
		word = strings.ToLower(word)
		processedWord := reg.ReplaceAllString(word, "")
		wordcounts[processedWord]++
	}

	if input.Err() != nil {
		log.Fatal(input.Err())
	}

	fmt.Printf("word\tfreq\n")

	for k, v := range wordcounts {
		fmt.Printf("%-12q\t%-5v\n", k, v)
	}
}
