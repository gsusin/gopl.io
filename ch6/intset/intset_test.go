// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercício 11.7

package intset

import (
    "fmt"
    "math"
    "math/rand"
    "gopl.io/ch11/intsetmap"
    "testing"
    "time"
)

func newRng() *rand.Rand {
    seed := time.Now().UTC().UnixNano()
    return rand.New(rand.NewSource(seed))
}

func BenchmarkAdd(b *testing.B) {
    var x IntSet
    rng := newRng()
    for i := 0; i < b.N; i++ {
        x.Add(rng.Intn(math.MaxUint32 / 2))
    }   
}

func BenchmarkAddMap(b *testing.B) {
    var x intsetmap.IntSetMap
    x.Init()
    rng := newRng()
    for i := 0; i < b.N; i++ {
        x.Add(rng.Intn(math.MaxUint32 / 2))
    }   
}

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}
