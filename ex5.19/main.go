package main

import "fmt"

func main() {
	fmt.Println(nonZeroWithoutReturn())
}

func nonZeroWithoutReturn() (val int) {
	defer func() {
		recover()
		val = 123
	}()
	panic("uhh")
}
