// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Modificado por Giancarlo Susin
// Exercícios 8.12, 8.13, 8.15

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster
type client struct {
	c    chan<- string // an outgoing message channel
	who  string
	conn net.Conn
}

type message struct {
	text string
	who  string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan message) // all incoming client messages
)

func broadcaster() {
	lastTime := make(map[string]time.Time)
	clients := make(map[client]bool) // all connected clients
	var tick <-chan time.Time
	tick = time.Tick(500 * time.Millisecond)
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				// Non blocking sending (couldn't test)
				select {
				case cli.c <- msg.text:
				default:
				}
			}
			lastTime[msg.who] = time.Now()
		case cli := <-entering:
			clients[cli] = true
			var names string
			for cli := range clients {
				if names != "" {
					names += ", "
				}
				names += cli.who
			}
			for cli := range clients {
				cli.c <- "current clients: " + names
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.c)
		case <-tick:
			//verifica o timeout de cada cliente e desconecta os que excederam
			for cli := range clients {
				if time.Now().After(lastTime[cli.who].Add(1 * time.Minute)) {
					for c := range clients {
						c.c <- cli.who + "closed by timeout"
					}
					delete(clients, cli)
					close(cli.c)
					cli.conn.Close()
				}
			}
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string, 3) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- message{who + " has arrived", who}
	entering <- client{ch, who, conn}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		s := who + ": " + input.Text()
		messages <- message{s, who}
	}
	if input.Err() != nil {
		// A conexão foi fechada por timeout, portanto os canais e a conexão já estão fechados
		return
	}
	leaving <- client{ch, who, conn}
	messages <- message{who + " has left", who}
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
