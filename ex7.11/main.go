// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

//!+main

func main() {
	prices := priceDatabase{db: map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/list", prices.list)
	http.HandleFunc("/price", prices.price)
	http.HandleFunc("/delete", prices.delete)
	http.HandleFunc("/update", prices.update)
	http.HandleFunc("/create", prices.create)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type priceDatabase struct {
	sync.Mutex
	db map[string]dollars
}

func (p priceDatabase) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range p.db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (p priceDatabase) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := p.db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		p.Lock()
		fmt.Fprintf(w, "no such item: %q\n", item)
		p.Unlock()
	}
}

func (p priceDatabase) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := p.db[item]; ok {
		p.Lock()
		delete(p.db, item)
		p.Unlock()
		fmt.Fprintf(w, "Deleted %s\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
func (p priceDatabase) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	oldPrice, ok := p.db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return

	}

	newPrice := req.URL.Query().Get("price")
	parsedNewPrice, err := strconv.ParseFloat(newPrice, 32)
	dollarPrice := dollars(parsedNewPrice)
	if err != nil {
		fmt.Fprintf(w, "Not a valid price: %s. Error: %v", newPrice, err)
		return
	}

	p.Lock()
	p.db[item] = dollarPrice
	fmt.Fprintf(w, "Updated %s from %s to %s", item, oldPrice, dollarPrice)
	p.Unlock()
}
func (p priceDatabase) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := p.db[item]; ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "Item already exists: %q\n", item)
		return
	}

	price := req.URL.Query().Get("price")
	intPrice, err := strconv.ParseFloat(price, 32)
	dollarPrice := dollars(intPrice)
	if err != nil {
		fmt.Fprintf(w, "Not a valid price: %s. Error: %v", price, err)
		return
	}

	p.Lock()
	p.db[item] = dollarPrice
	fmt.Fprintf(w, "Added entry: %s - %s", item, dollarPrice)
	p.Unlock()

}
