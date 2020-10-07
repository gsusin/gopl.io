// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Modificado por Giancarlo Susin
// Exercícios 2.4, 2.5

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

var pc [256]byte

//func init() {
//	doInit()
//}

func doInit() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	if pc[1] == 0 {
		doInit()
	}
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountAllBits(x uint64) int {
	var c int
	for i := 0; i < 64; i++ {
		c += int(x & 1)
		x >>= 1
	}
	return c
}

func PopCountCleanBit(x uint64) int {
	var c int
	for x > 0 {
		c++
		x = x & (x - 1)
	}
	return c
}

//!-
