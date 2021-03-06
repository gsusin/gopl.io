// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Modificado por Giancarlo Susin
// Exercício 7.11

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

var mu sync.Mutex

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	if _, ok := db[item]; ok {
		s := req.URL.Query().Get("price")
		if price, err := strconv.ParseFloat(s, 32); err == nil {
			db[item] = dollars(price)
			fmt.Fprintf(w, "price updated\n")
		} else {
			fmt.Fprintf(w, "invalid price: %s\n", s)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	mu.Unlock()
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	if _, ok := db[item]; ok {
		fmt.Fprintf(w, "item %s exists\n", item)
	} else {
		s := req.URL.Query().Get("price")
		if price, err := strconv.ParseFloat(s, 32); err == nil {
			db[item] = dollars(price)
			fmt.Fprintf(w, "item created\n")
		} else {
			fmt.Fprintf(w, "invalid price: %s\n", s)
		}
	}
	mu.Unlock()
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "element %s deleted\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %s\n")
	}
	mu.Unlock()
}
