package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
)

var tokens = make(chan struct{}, 20)
var maxDepth int
var seen = make(map[string]bool)
var sl = sync.Mutex{}

func main() {
	flag.IntVar(&maxDepth, "depth", 3, "max crawl depth")
	flag.Parse()
	wg := &sync.WaitGroup{}
	for _, link := range flag.Args() {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			crawl(link, 0, wg)
		}(link)
	}
	wg.Wait()
}

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(depth, url)
	if depth >= maxDepth {
		return
	}
	tokens <- struct{}{} // acquire token
	list, err := Extract(url)
	<-tokens // release token
	if err != nil {
		log.Print(err)
	}
	for _, link := range list {
		sl.Lock()
		if seen[link] {
			sl.Unlock()
			continue
		}
		seen[link] = true
		sl.Unlock()
		wg.Add(1)
		go crawl(link, depth+1, wg)
	}
}
