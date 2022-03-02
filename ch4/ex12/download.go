package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var index ComicIndex
	baseUrl := "https://xkcd.com/"
	suffix := "/info.0.json"
	endComic := 2540
	for i := 1; i <= endComic; i++ {
		comicURL := baseUrl + strconv.Itoa(i) + suffix
		fmt.Println(i)
		resp, err := http.Get(comicURL)

		if err != nil {
			log.Print(err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			log.Printf("Fetch failed for comic #%v: %s", i, resp.Status)
			continue
		}

		var current Comic
		if err := json.NewDecoder(resp.Body).Decode(&current); err != nil {
			log.Printf("JSON Unmarshaling Failed: %s", err)
			resp.Body.Close()
			continue
		}

		current.Link = baseUrl + strconv.Itoa(i)
		index.Comics = append(index.Comics, current)
		resp.Body.Close()
	}

	fmt.Printf("%+v", index.Comics)
	file, _ := json.Marshal(index)
	_ = ioutil.WriteFile("comics.json", file, 0644)
}
