package main

type Comic struct {
	Number     int    `json:"num"`
	Transcript string `json:"transcript"`
	Title      string `json:"title"`
	Link       string
}

type ComicIndex struct {
	Comics []Comic
}
