package popcount

import (
	"testing"
)

func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(123)
	}
}

func BenchmarkPopcountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(123)
	}
}
