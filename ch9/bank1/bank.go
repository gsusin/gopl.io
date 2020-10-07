// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Modificado por Giancarlo Susin
// Exercício 9.1

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

type withdraw struct {
	a     int
	rchan *chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan withdraw)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	rchan := make(chan bool)
	withdraws <- withdraw{amount, &rchan}
	return <-rchan
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdraws:
			switch w.a <= balance {
			case true:
				balance -= w.a
				*w.rchan <- true
			case false:
				*w.rchan <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
