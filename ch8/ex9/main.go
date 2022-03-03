package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type DuData struct {
	root   string
	c      chan int64
	nbytes int64
	nfiles int64
}

var vFlag = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// ...determine roots...

	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	var n sync.WaitGroup

	DuList := []DuData{}

	for _, root := range roots {
		ch := make(chan int64)
		du := DuData{root: root, c: ch}
		DuList = append(DuList, du)
		n.Add(1)
		go walkDir(du.root, &n, &du)
	}

	go func() {
		n.Wait()
		for _, du := range DuList {
			close(du.c)
		}
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
loop:
	for {
		for i := range DuList {
			select {
			case size, ok := <-DuList[i].c:
				if !ok {
					break loop // fileSizes was closed
				}
				DuList[i].nfiles++
				DuList[i].nbytes += size
			case <-tick:
				printDiskUsage(DuList)
			}
		}
	}

	printDiskUsage(DuList) // final totals
	// ...select loop...
}

func printDiskUsage(dus []DuData) {
	for _, du := range dus {
		fmt.Printf("%s: %d files %.9fGB\t", du.root, du.nfiles, float64(du.nbytes)/1e9)
	}
	fmt.Print("\n")
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, du *DuData) {
	//fmt.Printf("Walking %s\n", dir)
	//fmt.Println(*du)
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, du)
		} else {
			//fmt.Printf("Trying to add %d bytes\n", entry.Size())
			du.c <- entry.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
