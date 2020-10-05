// Exercício 7.10

package main

import (
	"fmt"
	"sort"
)

type sequence []rune

func (s sequence) Len() int { return len(s) }

func (s sequence) Less(i, j int) bool {
	//fmt.Printf("i=%d s[i]=%v j=%d s[j]=%v\n", i, s[i], j, s[j])
	return s[i] < s[j]
}

func (s sequence) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func IsPalindrome(s sort.Interface) bool {
	var j int
	for i := 0; i < s.Len(); i++ {
		j = s.Len() - 1 - i
		if !(s.Less(i, j) || !s.Less(j, i)) {
			return false
		}
	}
	return true
}

func main() {
	seq := sequence{'m', 'a', 'i', 'a', 'm'}
	fmt.Printf("É um palíndromo? %v\n", IsPalindrome(seq))
}
