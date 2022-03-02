// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
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

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
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

//!-breadthFirst

//!+crawl
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

//!-crawl

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

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	rootDir = os.Args[1]
	breadthFirst(crawl, os.Args[1:])
}

//!-main
