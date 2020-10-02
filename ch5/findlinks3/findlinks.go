// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercício 5.13
// Modificado por Giancarlo Susin em 09/06/2020

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopl.io/ch5/links"
)

var originalDomains map[string]bool

var dirTopo string

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	originalDomains = make(map[string]bool)
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				originalDomains[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, conteúdo, err := links.Extract2(url)
	if err != nil {
		log.Print(err)
	}
	for _, val := range list {
		writePage(val, conteúdo)
	}
	return list
}

func writePage(page string, conteúdo []byte) {
	if originalDomains[page] {
		var perm os.FileMode = 0770
		ajustar := func(r rune) rune {
			switch r {
			case '/':
				return '-'
			case ':':
				return '-'
			}
			return r
		}
		nome := strings.Map(ajustar, page)

		if os.Chdir(dirTopo) != nil {
			log.Print("Erro em chdir dirTopo.")
			return
		}
		if err := os.Mkdir(nome, perm); err != nil && !os.IsExist(err) {
			log.Print("Erro em mkdir.")
			return
		}
		if os.Chdir(nome) != nil {
			log.Print("Erro em chdir.")
			return
		}
		arq, err := os.Create(nome + ".html")
		if err != nil {
			log.Print("Erro em create.")
			return
		}
		_, errw := arq.Write(conteúdo)
		if errw != nil {
			log.Print("Erro em write.")
			return
		}

	}
}

//!-crawl

//!+main
func main() {
	//Poderia ter sido utilizada Getwd()
	dirTopo = os.Getenv("PWD")

	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

//!-main
