package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopl/ch4/ex12/comics"
)

func main() {

	//Load Comics Index
	comicFile, err := os.Open("comics.json")
	if err != nil {
		log.Fatal(err)
	}
	defer comicFile.Close()

	byteValue, _ := ioutil.ReadAll(comicFile)

	var XKCD comics.ComicIndex

	json.Unmarshal(byteValue, &XKCD)

	//Search Comics Index for Matches
	if len(os.Args) > 2 || len(os.Args) < 2 {
		log.Fatal("Usage: provide one search term per query")
	}
	search := strings.ToLower(os.Args[1])
	for _, e := range XKCD.Comics {
		title := strings.ToLower(e.Title)
		body := strings.ToLower(e.Transcript)
		if strings.Contains(title, search) || strings.Contains(body, search) {
			fmt.Printf("\nCartoon #%v: %v\nLink: %v\n\n%v\n", e.Number, e.Title, e.Link, e.Transcript)
		}
	}

}
