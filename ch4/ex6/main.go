package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	strings := []string{
		"a b",
		"a  b",
		"a   b",
		"a b c",
		"a      b      c",
		"Hello \u03C9\n\t\u0085\u00a0\x64",
	}
	for _, s := range strings {
		sq := string(squashSpace([]byte(s)))
		fmt.Printf("%q -> %q\n", s, sq)
	}
}

func squashSpace(bytes []byte) []byte {
	j := 0
	space := false

	for i, r := range string(bytes) {
		if unicode.IsSpace(r) {
			space = true
			continue
		} else if space {
			space = false
			bytes[j] = ' '
			j++
		}

		rl := utf8.RuneLen(r)
		copy(bytes[j:], bytes[i:i+rl])
		j += rl
	}
	return bytes[:j]
}
