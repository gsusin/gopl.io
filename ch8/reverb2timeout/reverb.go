// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Modificado por Giancarlo Susin
// Exercício 8.8

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	ch := make(chan string)
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			ch <- input.Text()
		}
	}()
	for {
		select {
		case msg := <-ch:
			go echo(c, msg, 1*time.Second)
		case <-time.After(10 * time.Second):
			// NOTE: ignoring potential errors from input.Err()
			c.Close()
			return
		}
	}
}

//func scanConn() {

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8003")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
