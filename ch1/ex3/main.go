package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func EchoNaive(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(ioutil.Discard, s)
}

func EchoBetter(args []string) {
	s := strings.Join(args, " ")
	fmt.Fprintln(ioutil.Discard, s)
}
