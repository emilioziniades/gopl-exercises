package main

import (
	"fmt"
)

func main() {
	fmt.Println("2 routines")
	makepipeline(2)

	fmt.Println("3 routines")
	makepipeline(3)

	fmt.Println("10 routines")
	makepipeline(10)
}

func startpipe(out chan<- int, n int) {
	for i := 0; i < n; i++ {
		out <- i
	}
	close(out)
}

func midpipe(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i
	}
	close(out)
}

func endpipe(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func makechans(npipes int) []chan int {
	chans := make([]chan int, npipes)
	for i := range chans {
		chans[i] = make(chan int)
	}
	return chans
}
func makepipeline(nroutines int) {
	if nroutines < 2 {
		return
	}
	npipes := nroutines - 1
	chans := makechans(npipes)

	//startpipe uses first chan
	go startpipe(chans[0], 1)

	//midpipe uses all pipes including  start and end,
	// from chans[0] until chans[npipes - 1]
	for i := 0; i < npipes-1; i++ {
		go midpipe(chans[i], chans[i+1])
	}
	//endpip uses only last chan, i.e.chans[npipes - 1]
	endpipe(chans[len(chans)-1])
}
