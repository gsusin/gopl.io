// Giancarlo Susin
// Exercício 11.6

package popcount

import (
	"testing"
)

var tests = []struct {
	val    uint64
	wanted int
}{
	{1, 1},
	{2, 1},
	{4, 1},
	{1073741823, 30},
	{0, 0},
}

func TestPopCountAllBits(t *testing.T) {
	for _, test := range tests {
		if v := PopCountAllBits(test.val); v != test.wanted {
			t.Errorf("PopCountAllBits(%v) = %d, want %d", test.val, v, test.wanted)
		}
	}
}

func TestPopCountCleanBit(t *testing.T) {
	for _, test := range tests {
		if v := PopCountCleanBit(test.val); v != test.wanted {
			t.Errorf("PopCountCleanBit(%v) = %d, want %d", test.val, v, test.wanted)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for _, test := range tests {
		for i := 0; i < b.N; i++ {
			PopCount(test.val)
		}
	}
}

func BenchmarkPopCountiCleanBit(b *testing.B) {
	for _, test := range tests {
		for i := 0; i < b.N; i++ {
			PopCountCleanBit(test.val)
		}
	}
}

func BenchmarkPopCountAllBits(b *testing.B) {
	for _, test := range tests {
		for i := 0; i < b.N; i++ {
			PopCountAllBits(test.val)
		}
	}
}
