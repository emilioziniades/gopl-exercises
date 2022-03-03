package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	Echo(os.Args[0:])
}

func Echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}
