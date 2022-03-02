package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
)

var trackTable = template.Must(template.New("tracktable").Parse(`
	<style>
		table, th, td {
			 border: 1px solid black;
		}
	</style>
	<h1> Sortable Music Table </h1>
	<table>
	<tr>
		<th><a href="?sort=title">Title</a></th>
		<th><a href="?sort=artist">Artist</a></th>
		<th><a href="?sort=album">Album</a></th>
		<th><a href="?sort=year">Year</a></th>
		<th><a href="?sort=length">Length</a></th>
	</tr>
	{{range .T}}
	<tr>
		<td>{{.Title}}</td>
		<td>{{.Artist}}</td>
		<td>{{.Album}}</td>
		<td>{{.Year}}</td>
		<td>{{.Length}}</td>
	</tr>
	{{end}}
	</table>
`))

func main() {
	table := CustomSort{tracks, []sortFunc{}}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.FormValue("sort") {
		case "title":
			table.addFunc(titleSort)
		case "artist":
			table.addFunc(artistSort)
		case "album":
			table.addFunc(albumSort)
		case "year":
			table.addFunc(yearSort)
		case "length":
			table.addFunc(lengthSort)
		}
		fmt.Println(table.C)
		sort.Sort(table)
		if err := trackTable.Execute(w, table); err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
