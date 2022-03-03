package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var rootDir string

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	err := download(url)
	if err != nil {
		panic(err)
	}

	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}

func download(s string) (err error) {

	u, err := url.Parse(s)
	if err != nil {
		return
	}

	if currentRoot := "https://" + u.Host; currentRoot != rootDir {
		fmt.Println("\tNot in rootdir, skipping:", s)
		return
	}

	resp, err := http.Get(s)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	filepath := "./" + u.Host + strings.TrimRight(u.Path, "/")
	filename := filepath + ".html"
	os.MkdirAll(filepath, 0700)

	fmt.Println("\tDownloading", s, "in", filename)

	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return

}

func main() {
	rootDir = os.Args[1]
	breadthFirst(crawl, os.Args[1:])
}
