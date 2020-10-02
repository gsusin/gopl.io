// Exercício 5.16 

package main

import "fmt"

func Join(sep string, elems ...string) string {
	j := ""
	sepdin := ""
	for _, e := range elems {
		j += sepdin + e
		sepdin = sep
	}
	return j
}

func main() {
	fmt.Printf("Join(*, g, i, a, n) = %s\n", Join("*", "g", "i", "a", "n"))
}
