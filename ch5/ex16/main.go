package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(Join(";", "a", "b", "c"))
	fmt.Println(Join(";", "a", "b"))
	fmt.Println(Join(";", "a"))
	fmt.Println(Join(";"))
}

func Join(sep string, elems ...string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}
	res := elems[0]
	for _, elem := range elems[1:] {
		res += sep
		res += elem
	}
	return res
}

func JoinTorbiak(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	b := bytes.Buffer{}
	for _, s := range strs[:len(strs)-1] {
		b.WriteString(s)
		b.WriteString(sep)
	}
	b.WriteString(strs[len(strs)-1])
	return b.String()
}
