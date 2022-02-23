package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Counting element types: %v\n", err)
		os.Exit(1)
	}
	for elem, count := range countElems(make(map[string]int), doc) {
		fmt.Printf("%v\t\t%v\n", elem, count)
	}
}

func countElems(elemCount map[string]int, n *html.Node) map[string]int {

	if n.Type == html.ElementNode {
		elemCount[n.Data]++
	}

	if c := n.FirstChild; c != nil {
		elemCount = countElems(elemCount, c)
	}

	if c := n.NextSibling; c != nil {
		elemCount = countElems(elemCount, c)
	}

	return elemCount
}
