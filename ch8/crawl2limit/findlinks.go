// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Modificado por Giancarlo Susin
// Exercício 8.6

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

type visited struct {
	links []string
	depth int
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

//!+
func main() {
	var maxDepth = flag.Int("depth", 1, "Informe a profundidade da busca")
	worklist := make(chan visited)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- visited{os.Args[1:], 0} }()

	// Crawl the web concurrently.
	// Each "seen" map value represents number of links traversed
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		if list.depth <= *maxDepth {
			for _, link := range list.links {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string, depth int) {
						worklist <- visited{crawl(link), depth}
					}(link, list.depth+1)
				}
			}
		}
	}
}

//!-
