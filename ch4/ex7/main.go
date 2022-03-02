package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	tests := [][]byte{
		[]byte("Hello"),
		[]byte("Володимир"),
	}
	for _, t := range tests {
		fmt.Println(string(t))
		reverseRunes(t)
		fmt.Println(string(t))
	}
}

func reverseRunes(bytes []byte) {
	l := len(bytes)
	pos := l

	// iterate over bytes forward, appending to back of bytes
	for _, r := range string(bytes) {
		pos -= utf8.RuneLen(r)
		copy(bytes[pos:], []byte(string(r)))

	}
}
