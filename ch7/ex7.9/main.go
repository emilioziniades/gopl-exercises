package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//*****************************************************************************

type comparison int

const (
	lt comparison = iota
	eq
	gt
)

type CustomSort struct {
	T []*Track
	C []sortFunc
}

func (x *CustomSort) addFunc(f sortFunc) { x.C = append(x.C, f) }

func (x CustomSort) Len() int           { return len(x.T) }
func (x CustomSort) Swap(i, j int)      { x.T[i], x.T[j] = x.T[j], x.T[i] }
func (x CustomSort) Less(i, j int) bool { return multiSort(x.C)(x.T[i], x.T[j]) }

func multiSort(sortFuncs []sortFunc) func(x, y *Track) bool {

	sort := func(x, y *Track) bool {
		for _, f := range sortFuncs {
			cmp := f(x, y)
			switch cmp {
			case eq:
				continue
			case lt:
				return true
			case gt:
				return false
			}
		}
		return false
	}
	return sort

}

type sortFunc func(x, y *Track) comparison

func artistSort(x, y *Track) comparison {
	switch {
	case x.Artist < y.Artist:
		return lt
	case x.Artist == y.Artist:
		return eq
	default:
		return gt
	}
}

func titleSort(x, y *Track) comparison {
	switch {
	case x.Title < y.Title:
		return lt
	case x.Title == y.Title:
		return eq
	default:
		return gt
	}
}

func albumSort(x, y *Track) comparison {
	switch {
	case x.Album < y.Album:
		return lt
	case x.Album == y.Album:
		return eq
	default:
		return gt
	}
}

func yearSort(x, y *Track) comparison {
	switch {
	case x.Year < y.Year:
		return lt
	case x.Year == y.Year:
		return eq
	default:
		return gt
	}
}

func lengthSort(x, y *Track) comparison {
	switch {
	case x.Length < y.Length:
		return lt
	case x.Length == y.Length:
		return eq
	default:
		return gt
	}
}

//func main() {
//	printTracks(tracks)
//	table := CustomSort{tracks, []sortFunc{}}
//
//	fmt.Println("\nClick on title\n")
//	table.addFunc(titleSort)
//	sort.Sort(table)
//	printTracks(tracks)
//
//	fmt.Println("\nClick on year\n")
//	table.addFunc(yearSort)
//	sort.Sort(table)
//	printTracks(tracks)
//
//}
