// Exercício 7.5

package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if _, err := io.Copy(os.Stdout, LimitReader(strings.NewReader("Giancarlo"), 4)); err != nil {
		log.Fatal(err)
	}
}

type LimitedReader struct {
	i int64
	r func([]byte) (int, error)
}

func (l LimitedReader) Read(p []byte) (n int, err error) {
	pl := p[:int(l.i)]
	n, err = l.r(pl)
	if err != nil {
		return n, err
	}
	if n < len(p) {
		return n, io.EOF
	}
	return n, nil
}

func LimitReader(r io.Reader, n int64) io.Reader {
	var lr LimitedReader
	lr.i = n
	lr.r = r.Read
	return &lr
}
