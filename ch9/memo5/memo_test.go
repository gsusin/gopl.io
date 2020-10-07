// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercício 9.3

package memo_test

import (
	"testing"
	"time"

	memo "gopl.io/ch9/memo5"
	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	done := make(chan struct{})
	m := memo.New(httpGetBody)
	defer m.Close()
	go func() {
		time.Sleep(6 * time.Second)
		close(done)
	}()
	memotest.Sequential(t, m, done)
}

func TestConcurrent(t *testing.T) {
	done := make(chan struct{})
	m := memo.New(httpGetBody)
	defer m.Close()
	go func() {
		//os.Stdin.Read(make([]byte, 1))
		time.Sleep(2 * time.Second)
		close(done)
	}()
	memotest.Concurrent(t, m, done)
}
