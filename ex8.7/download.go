package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {

	link := os.Args[1]
	download(link)
}

func download(link string) (err error) {

	fmt.Println(link)

	// Get webpage
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("getting %s: %s", link, resp.Status)
	}
	defer resp.Body.Close()

	//Extract and change links
	//doc, err := html.Parse(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}

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
	//ht, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("%s \n", ht)
	if _, err := io.Copy(f, resp.Body); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	return nil
}
