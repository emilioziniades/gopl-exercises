package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("enter an expression:")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\t")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	expr, err := Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	varlist := make(map[Var]bool)
	err = expr.Check(varlist)
	if err != nil {
		log.Println(err)
	}

	vars := make(Env)
	for v := range varlist {
		fmt.Printf("\t%s = ", v)
		in, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err)
		}
		in = strings.TrimRight(in, "\n")
		inFloat, err := strconv.ParseFloat(in, 64)
		if err != nil {
			log.Printf("Cant convert %q to float. Using 0", in)
		}
		vars[v] = inFloat
	}

	fmt.Printf("%s = %v\n", expr, expr.Eval(vars))
}
