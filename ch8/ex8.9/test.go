package main

import (
	"fmt"
	"time"
)

type chanStruct struct {
	c     chan int
	count int
}

func send(cs *chanStruct) {
	cs.c <- 42
}

func recv(cs *chanStruct) {
	i, ok := <-cs.c
	if !ok {
		fmt.Println("Not ok")
	}
	cs.count += i
	fmt.Println(i)
	fmt.Println(cs.count)
}
func main() {
	cs := chanStruct{c: make(chan int)}

	go send(&cs)
	go recv(&cs)

	defer func() {
		fmt.Println(cs)
	}()

	time.Sleep(3 * time.Second)
}
