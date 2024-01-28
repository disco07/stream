package stack

import "testing"

func BenchmarkStack_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := New[int]()
		for j := 0; j < 1000; j++ {
			stack.Push(j)
		}
	}
}
