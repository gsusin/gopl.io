// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Modificado por Giancarlo Susin
// Exercício 8.11

// See page 148.

// Fetch saves the contents of a URL into a local file.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

//!+
// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) {
	resp, err := http.Get(url)
	if cancelled() {
		return
	}
	if err != nil {
		responses <- response{url, "", 0, err}
		return
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		responses <- response{url, "", 0, err}
		return
	}
	n, err := io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	responses <- response{url, local, n, err}
}

//!-
type response struct {
	url      string
	filename string
	n        int64
	err      error
}

var responses = make(chan response, 20)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false

	}
}

func main() {
	for _, url := range os.Args[1:] {
		go fetch(url)
	}
	r := <-responses
	close(done)
	url, local, n, err := r.url, r.filename, r.n, r.err
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
	}
	fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
}
