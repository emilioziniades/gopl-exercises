// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	Echo(os.Args[0:])
}

func Echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

//!-