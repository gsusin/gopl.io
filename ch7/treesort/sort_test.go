// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"gopl.io/ch7/treesort"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func ExampleString() {
	data := make([]int, 3)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	fmt.Printf("Valores da árvore:\n%s", treesort.LastTree)
	// Output:
	// Valores da árvore:
	// 5
	//   L
	//   R
	//   21
	//     L
	//     17
	//       L
	//       R
	//     R
}
