// Exercício 11.1

package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	var b strings.Builder
	b.WriteString("Aproximação: 32")
	b.Write([]byte{0xff, 0xff, 0xff, 0xff})
	s := b.String()
	in := bufio.NewReader(strings.NewReader(s))
	c := count(in)
	if c.counts["letter"] != 11 {
		t.Errorf("count(%s).counts[\"letter\"] = %d, want %d", s, c.counts["letter"], 11)
	}
	if c.counts["digit"] != 2 {
		t.Errorf("count(%s).counts[\"digit\"] = %d, want %d", s, c.counts["digit"], 2)
	}
	if c.counts["space"] != 1 {
		t.Errorf("count(%s).counts[\"space\"] = %d, want %d", s, c.counts["space"], 1)
	}
	if c.counts["others"] != 1 {
		t.Errorf("count(%s).counts[\"others\"] = %d, want %d", s, c.counts["others"], 0)
	}
	if c.invalid != 4 {
		t.Errorf("count(%s).invalid = %d, want %d", s, c.invalid, 4)
	}
}
