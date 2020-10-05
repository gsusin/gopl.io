// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercício 7.1

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

//!+bytecounter

type ByteCounter int

type WordCounter int

type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	cont := 0
	for scanner.Scan() {
		cont++
	}
	*c += WordCounter(cont)
	var err error
	if err = scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return cont, err
}

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	cont := 0
	for scanner.Scan() {
		cont++
	}
	*c += LineCounter(cont)
	var err error
	if err = scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return cont, err
}

//!-bytecounter

func main() {
	//!+main
	var c ByteCounter
	var w WordCounter
	var l LineCounter
	c.Write([]byte("hello florianópolis SC"))
	w.Write([]byte("hello florianópolis SC\nBrasil"))
	l.Write([]byte("hello florianópolis SC\nBrasil"))
	fmt.Println(c) // "5", = len("hello")
	fmt.Println(w)
	fmt.Println(l)

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main
}
