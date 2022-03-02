package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	for _, e := range os.Args[1:] {
		ss := strings.Split(e, ",")
		fmt.Printf("%s anagram of %s? %v\n", ss[0], ss[1], isAnagram(ss[0], ss[1]))
	}
}

func isAnagram(s1, s2 string) bool {
	letters1 := make(map[rune]int)
	letters2 := make(map[rune]int)

	for _, letter := range s1 {
		letters1[letter]++
	}

	for _, letter := range s2 {
		letters2[letter]++
	}

	return reflect.DeepEqual(letters1, letters2)
}
