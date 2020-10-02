// Exercício 4.9

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "lendo a entrada:", err)
	}
	for palavra, frequência := range counts {
		fmt.Printf("%s\t%d\n", palavra, frequência)
	}

}
