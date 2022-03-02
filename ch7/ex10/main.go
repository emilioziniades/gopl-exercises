package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := sort.IntSlice([]int{1, 2, 3, 2, 1})
	r := sort.IntSlice([]int{1, 2, 3, 3, 2, 1})
	t := sort.StringSlice(stringToSlice("Hello World"))
	u := sort.StringSlice(stringToSlice("neveroddoreven"))
	fmt.Println(IsPalindrome(s))
	fmt.Println(IsPalindrome(r))
	fmt.Println(IsPalindrome(t))
	fmt.Println(IsPalindrome(u))
}

func IsPalindrome(s sort.Interface) bool {
	n := s.Len() - 1
	m := s.Len() / 2
	for i := 0; i < m; i++ {
		if s.Less(i, n-i) || s.Less(n-i, i) {
			return false
		}
	}
	return true
}

func stringToSlice(s string) []string { return strings.Split(s, "") }
