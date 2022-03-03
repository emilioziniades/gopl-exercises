package main

import (
	"fmt"
	"os"
)

func main() {
	Echo(os.Args[0:])
}

func Echo(args []string) {
	for i, e := range args {
		fmt.Printf("%v : %v\n", i, e)
	}
}
