package intsetmap

// Giancarlo Susin
// 31/07/2020

import (
	"fmt"
	"sort"
)

type IntSetMap struct {
	words map[int]bool
}

func (s *IntSetMap) Init() {
	s.words = make(map[int]bool)
}

func (s *IntSetMap) Has(x int) bool {
	return s.words[x]
}

func (s *IntSetMap) Add(x int) {
	s.words[x] = true
}

func (s *IntSetMap) UnionWith(t *IntSetMap) {
	for k := range t.words {
		(*s).words[k] = true
	}
}

func (s *IntSetMap) String() string {
	var out string
	var ints []int

	for i := range (*s).words {
		ints = append(ints, i)
	}
	sort.Ints(ints)

	for _, k := range ints {
		if out == "" {
			out = fmt.Sprintf("{%d", k)
			continue
		}
		out += fmt.Sprintf(" %d", k)
	}
	if out == "" {
		return "{}"
	}
	return out + "}"
}
