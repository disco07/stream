package deque

import "testing"

func BenchmarkDequePushBack(b *testing.B) {
	deque := New[int]()
	for i := 0; i < b.N; i++ {
		deque.PushBack(i)
	}

	for i := 0; i < b.N; i++ {
		if deque.data[i] != i {
			b.Errorf("deque.data[%d] != %d", deque.data[i], i)
		}
	}
}

func BenchmarkSliceAppend(b *testing.B) {
	var slice []int
	for i := 0; i < b.N; i++ {
		slice = append(slice, i)
	}

	for i := 0; i < b.N; i++ {
		if slice[i] != i {
			b.Errorf("slice[%d] != %d", slice[i], i)
		}
	}
}
