package unordered_map

import (
	"testing"
)

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		for j := 0; j < 1000; j++ {
			m[j] = j
		}
	}
}

func BenchmarkUnorderedMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := New[int, int](Hash[int])
		for j := 0; j < 1000; j++ {
			m.Set(j, j)
		}
	}
}
