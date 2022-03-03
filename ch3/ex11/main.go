package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	parts := strings.Split(s, ".")
	var whole, decimal string
	whole = parts[0]
	if len(parts) > 1 {
		decimal = parts[1]
	}

	var buf bytes.Buffer
	if whole[0] == '-' || s[0] == '+' {
		buf.WriteByte(whole[0])
		whole = whole[1:]
	}
	n := len(whole)
	c := (n + 2) % 3

	for i := 0; i < n; i++ {
		buf.WriteByte(whole[i])
		if i == n-1 {
			continue
		}
		if c == 0 {
			buf.WriteRune(',')
			c = 3
		}
		c--
	}

	if decimal != "" {
		buf.WriteRune('.')
		buf.WriteString(decimal)
	}
	return buf.String()
}
