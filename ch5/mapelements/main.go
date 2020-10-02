// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercício 5.2

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for e, q := range visit(nil, doc, nil) {
		fmt.Printf("%s\t%d\n", e, q)
	}
	fmt.Printf("%v\n", doc)
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(elementos map[string]int, n *html.Node, nextSibling *html.Node) map[string]int {
	if elementos == nil {
		elementos = make(map[string]int)
	}
	if n.Type == html.ElementNode {
		elementos[n.Data]++
	}
	if n.FirstChild != nil {
		elementos = visit(elementos, n.FirstChild, n.FirstChild.NextSibling)
	}
	if nextSibling != nil {
		elementos = visit(elementos, nextSibling, nextSibling.NextSibling)
	}
	return elementos
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
