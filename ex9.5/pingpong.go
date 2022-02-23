package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var n int
	go func() {
		ch1 <- 1
		for i := range ch2 {
			n++
			ch1 <- i
		}
	}()

	go func() {
		for i := range ch1 {
			ch2 <- i
		}
	}()

	<-time.After(10 * time.Second)
	fmt.Println(n)
	close(ch1)
	close(ch2)
}
