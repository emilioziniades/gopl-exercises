package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Package struct {
	Path string   `json:"ImportPath"`
	Name string   `json:"Name"`
	Deps []string `json:"Deps"`
}

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		log.Fatal("usage: specify any number of packages as arguments.")
	}

	all := getAllPackages()
	for _, e := range os.Args[1:] {
		format := "(%s) is transitively depended on by\n"
		fmt.Printf(format, e)
		initial := getPackage(e)
		initialDependents := make([]string, 0)
		for _, e := range all {
			if has(e.Deps, initial.Path) {
				initialDependents = append(initialDependents, e.Path)
			}
		}
		for _, e := range initialDependents {
			fmt.Printf("\t-> %s\n", e)
		}
	}
}

func getPackage(s string) (pkg Package) {
	output := goList(s)
	if err := json.Unmarshal(output, &pkg); err != nil {
		log.Fatal(err)
	}
	return
}

func getAllPackages() (pkgs []Package) {
	output := goList("...")
	decoder := json.NewDecoder(bytes.NewReader(output))
	for {
		var currPkg Package
		if err := decoder.Decode(&currPkg); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		pkgs = append(pkgs, currPkg)
	}
	return
}

func goList(arg string) []byte {
	out, err := exec.Command("go", "list", "-json", arg).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func has(slice []string, s string) bool {
	for _, e := range slice {
		if e == s {
			return true
		}
	}
	return false
}
