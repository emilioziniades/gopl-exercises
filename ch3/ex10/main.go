package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	c := (n + 2) % 3

	for i := 0; i < n; i++ {
		buf.WriteByte(s[i])
		if i == n-1 {
			continue
		}
		if c == 0 {
			buf.WriteRune(',')
			c = 3
		}
		c--
	}
	return buf.String()
}
