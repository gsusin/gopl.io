// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+

// Modificado por Giancarlo Susin
// Exercício 7.17

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // stack of element names
	var attribs []xml.Attr
	var attrSearch bool = strings.Contains(os.Args[1], "=\"")
	var attrName, attrValue string
	if attrSearch {
		parts := strings.Split(os.Args[1], "=\"")
		attrName = parts[0]
		attrValue = strings.TrimSuffix(parts[1], "\"")
	}
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
			attribs = tok.Attr
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if !attrSearch {
				if containsAll(stack, os.Args[1:]) {
					fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
				}
			} else {
				if containsAttr(attribs, attrName, attrValue) {
					fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
				}
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

// containsAttr report whether s contains attribute sa in any position
func containsAttr(s []xml.Attr, sn string, sv string) bool {
	for _, v := range s {
		if string(v.Name.Local) == sn && string(v.Value) == sv {
			return true
		}
	}
	return false
}

//!-
