// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercício 5.10

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	courses, nivel := topoSort(prereqs)
	for i := range courses {
		fmt.Printf("%d:\t%s %d\n", i+1, courses[i], nivel[i])
	}
}

func topoSort(m map[string][]string) ([]string, []int) {
	var order []string
	var nivel []int
	seen := make(map[string]bool)
	var visitAll func(items map[string][]string) int

	visitAll = func(items map[string][]string) int {
		ntotal := 1
		for key, _ := range items {
			//fmt.Printf("Passou por: %s\n", item)
			nsub := 1
			if !seen[key] {
				seen[key] = true
				prox, ok := m[key]
				if ok {
					filhos := make(map[string][]string)
					for _, filho := range prox {
						filhos[filho] = nil
					}
					nsub = visitAll(filhos) + 1
				}
				order = append(order, key)
				nivel = append(nivel, nsub)
			}
			if ntotal < nsub {
				ntotal = nsub
			}
		}
		return ntotal
	}
	visitAll(m)
	return order, nivel
}

//!-main
