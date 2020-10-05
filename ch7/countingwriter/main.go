// Exercício 7.2

package main

import (
	"fmt"
	"io"
	"os"
)

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var b ByteCounter
	b.w = w.Write
	return &b, &b.i
}

type ByteCounter struct {
	i int64
	w func([]byte) (int, error)
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	(*c).i += int64(len(p))
	(*c).w(p)
	return len(p), nil
}

func main() {
	w, c := CountingWriter(os.Stdout)
	n := "Giancarlo Susin"
	fmt.Fprintf(w, "Nome: %s\n", n)
	fmt.Fprintf(w, "Quantidade: %v\n", *c)
}
