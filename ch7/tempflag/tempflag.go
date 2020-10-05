// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 181.

// Exercício 7.6

// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"

	"gopl.io/ch7/tempconv"
)

//!+
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")
var tempK = tempconv.KelvinFlag("tempK", 300.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
	fmt.Println(*tempK)
}

//!-
