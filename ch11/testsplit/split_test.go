// Exercício 11.5

package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		inputString string
		inputSep    string
		want        int
	}{
		{"a:b:c", ":", 3},
		{"a:b:c", ",", 1},
		{"a b c", " ", 3},
	}
	for _, test := range tests {
		if got := strings.Split(test.inputString, test.inputSep); len(got) != test.want {
			t.Errorf("Split(%q, %q) = %d, want %d", test.inputString, test.inputSep, len(got), test.want)
		}
	}
}
