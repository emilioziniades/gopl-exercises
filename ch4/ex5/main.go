package main

import "fmt"

func main() {
	a := []int{1, 2, 2, 3, 3, 3, 4, 5}
	fmt.Println(a)
	a = dedup(a)
	fmt.Println(a)
}

func dedup(ints []int) []int {
	j := 0
	for i := 1; i < len(ints); i++ {
		if ints[j] == ints[i] {
			continue
		}
		j++
		ints[j] = ints[i]
	}
	ints = ints[:j+1]
	return ints

}
