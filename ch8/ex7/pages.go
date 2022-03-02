package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"

	"golang.org/x/net/html"
)

// Fetch web page
// Iterate over nodes and change links
// Save web pages in nested folders

var tokens = make(chan struct{}, 20)

func getPage(link string, wg *sync.WaitGroup) (links []string) {
	defer wg.Done()
	// Fetch web page
	fmt.Println(link)
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("getting %s: %s", link, resp.Status)
	}
	defer resp.Body.Close()

	//Extract and change links
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}

				fmt.Println(a.Val)
				// Add links to worklist
				links = append(links, link.String())

				//Update link to local file
				a.Val = "./" + link.Path + "/index.html"
				n.Attr[i] = a
				fmt.Println(a.Val)
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	// Save web page
	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}
	filepath := u.Host + u.Path
	fmt.Println("downloading", link, " to ", filepath)
	if err := os.MkdirAll(filepath, 0700); err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(filepath + "/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Render(f, doc)
	if err != nil {
		log.Fatal(err)
	}
	//	if _, err := io.Copy(f, resp.Body); err != nil {
	//		log.Fatal(err)
	//	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	return links
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var wg = &sync.WaitGroup{}

func main() {
	//parse first URL
	if len(os.Args) != 2 {
		log.Fatal("usage: ./mirror url")
	}
	root, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var n int // number of pending sends to worklist
	seen := make(map[string]bool)
	worklist := make(chan []string)

	n++
	go func() { worklist <- os.Args[1:] }()

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			u, err := url.Parse(link)
			if err != nil {
				continue
			}
			if !seen[link] && u.Host == root.Host {
				seen[link] = true
				n++
				wg.Add(1)
				go func(link string) {
					tokens <- struct{}{} // acquire a token
					worklist <- getPage(link, wg)
					<-tokens // release the token
				}(link)
			}
		}
	}
	wg.Wait()
}
