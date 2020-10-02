// Exercício 5.19

package main

import "fmt"

func noreturn(e int) (r int) {
	defer func() {
		if recover() != nil {
			r = 99
		}
	}()
	panic(e)
}

func main() {
	fmt.Printf("noreturn(7)=%d\n", noreturn(7))
}
