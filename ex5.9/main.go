package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(expand("Hey $there my $dude", upper))
}

func upper(s string) string {
	return strings.ToUpper(s)
}

// expand replaces each substring "$foo" with f("foo")
func expand(s string, f func(string) string) string {
	re, _ := regexp.Compile(`\$[a-zA-Z]+`)
	match := re.FindAll([]byte(s), -1)
	fmt.Printf("%q\n", match)

	res := s

	for _, e := range match {
		es := string(e)
		res = strings.ReplaceAll(res, es, f(es[1:]))
	}
	return res

}
