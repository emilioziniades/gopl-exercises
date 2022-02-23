package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := getDocument(os.Args[1])
	if err != nil {
		panic(err)
	}
	//	images := ElementsByTagName(doc, "img")
	headers := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	for _, e := range headers {
		fmt.Printf("%+v\n", *e)
	}
}

func getDocument(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return doc, nil
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var matches []*html.Node

	matchElement := func(n *html.Node, tag string) {
		if n.Data == tag {
			matches = append(matches, n)
		}
	}

	var forEachNode func(n *html.Node, tag string)
	forEachNode = func(n *html.Node, tag string) {

		matchElement(n, tag)

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c, tag)
		}
	}

	for _, e := range name {
		forEachNode(doc, e)
	}
	return matches
}
