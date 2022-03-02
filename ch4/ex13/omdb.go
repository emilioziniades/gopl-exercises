package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Movie struct {
	Title  string
	Poster string
}

func main() {
	title := parseInput()
	req := constructQuery(title)

	log.Printf("Fetching poster for %v", title)

	resp, err := http.Get(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("Query Error Status Code %v", resp.Status)
	}
	var movie Movie

	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		resp.Body.Close()
		log.Fatal(err)
	}

	if empty := new(Movie); *empty == movie {
		log.Fatalf("No result found for %v", title)
	}

	resp.Body.Close()

	fetchImage(movie.Poster, strings.ReplaceAll(title, " ", "_"))
}

func constructQuery(title string) string {
	URL, err := url.Parse("http://www.omdbapi.com/")
	if err != nil {
		log.Fatal(err)
	}
	query := URL.Query()
	query.Set("apikey", "8aedc204")
	query.Set("t", title)
	URL.RawQuery = query.Encode()
	return URL.String()
}

func parseInput() string {
	if len(os.Args) > 2 || len(os.Args) < 2 {
		log.Fatal("Specify a film title in quotation marks e.g. \"Iron Man\" ")
	}
	return os.Args[1]
}

func fetchImage(imgUrl, fileName string) {
	resp, err := http.Get(imgUrl)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}
	file, err := os.Create(fileName + ".jpg")
	_, err = io.Copy(file, resp.Body)

	log.Printf("Saving poster to %v.jpg", fileName)
	file.Close()
}
