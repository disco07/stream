package deque

import (
	"testing"
)

// Benchmark pour PushBack sur votre Deque
func BenchmarkDequePushBack(b *testing.B) {
	deque := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		deque.PushBack(i)
	}

	for i := 0; i < b.N; i++ {
		if deque.data[i] != i {
			b.Errorf("deque.data[%d] != %d", deque.data[i], i)
		}
	}
}

// Benchmark pour PushBack sur une slice
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

// Benchmark pour PushFront sur votre Deque
func BenchmarkDequePushFront(b *testing.B) {
	deque := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		deque.PushFront(i)
	}

	for i := 0; i < b.N; i++ {
		if deque.data[i] != deque.Size()-i {
			b.Errorf("deque.data[%d] != %d", deque.data[i], deque.Size())
		}
	}
}

//	func TestNewDequePushBack(t *testing.T) {
//		deque := NewDeque[int]()
//
//		for i := 0; i < 10000; i++ {
//			deque.PushBack(i)
//		}
//
//		iterator := deque.Iterator()
//		for iterator.Next() {
//			log.Println(iterator.Value())
//		}
//	}
//func TestNewDequePushFront(t *testing.T) {
//	deque := NewDeque[int]()
//
//	for i := 0; i < 1000000; i++ {
//		deque.PushFront(i)
//	}
//
//	iterator := deque.Iterator()
//	for iterator.Next() {
//		log.Println(iterator.Value())
//	}
//}
