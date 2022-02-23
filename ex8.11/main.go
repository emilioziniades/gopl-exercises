// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var cancel = make(chan struct{})

func main() {
	urls := os.Args[1:]
	response := make(chan string, len(urls))
	for _, url := range urls {
		go func(url string) { response <- fetch(url) }(url)
	}
	fmt.Println(<-response)
	close(cancel)
}

func fetch(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	return fmt.Sprintf("%s", b)
}

//!-
