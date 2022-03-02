package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "textnodeprint: %v\n", err)
		os.Exit(1)
	}

	printTextNodes(doc)
}

func printTextNodes(n *html.Node) {

	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	if c := n.FirstChild; c != nil && c.Data != "style" && c.Data != "script" {
		printTextNodes(c)
	}
	if s := n.NextSibling; s != nil && s.Data != "style" && s.Data != "script" {
		printTextNodes(s)
	}
}
