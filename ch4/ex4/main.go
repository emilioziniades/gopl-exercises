package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	rotate(a, 2)
	fmt.Println(a)
}

func rotate(s []int, i int) {
	copy(s, append(s[i:], s[:i]...))
}
