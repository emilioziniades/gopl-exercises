package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		w, i, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Printf("CountWordsAndImages: %v")
		}
		fmt.Printf("%v: %v words and %v images", url, w, i)
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		fmt.Println(n.Data)
		images++
	}
	if n.Type == html.TextNode && n.Data != "style" && n.Data != "script" {
		currentWords := strings.Fields(n.Data)
		fmt.Println(currentWords)
		words += len(currentWords)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		newWords, newImages := countWordsAndImages(c)
		words += newWords
		images += newImages
	}

	return
}
