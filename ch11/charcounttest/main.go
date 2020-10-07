// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type counting struct {
	counts  map[string]int
	invalid int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	c := count(in)

	fmt.Printf("rune\tcount\n")
	for c, n := range c.counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	if c.invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", c.invalid)
	}
}

func count(in *bufio.Reader) counting {
	c := counting{make(map[string]int), 0}
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			c.invalid++
			continue
		}
		switch {
		case unicode.IsLetter(r):
			c.counts["letter"]++
		case unicode.IsDigit(r):
			c.counts["digit"]++
		case unicode.IsSpace(r):
			c.counts["space"]++
		case unicode.IsNumber(r):
			c.counts["number"]++
		default:
			c.counts["others"]++

		}
	}
	return c
}

//!-
