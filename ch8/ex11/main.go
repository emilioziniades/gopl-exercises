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
