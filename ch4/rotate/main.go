// Exercício 4.4

package main

import "fmt"

func rotate(s []string, n int) []string {
	var t = make([]string, len(s))
	copy(t, s)
	n %= len(s)
	for i, j := len(s)-1, len(s)-1-n; i >= 0; i, j = i-1, (j-1+len(s))%len(s) {
		t[j] = s[i]
	}
	return t
}

func main() {
	data := [...]string{"7", "2", "3", "9", "0", "1"}
	fmt.Printf("%q\n%v\n", data, rotate(data[:], 2))
}
