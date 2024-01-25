package deque

type Iterator[T any] struct {
	deque *Deque[T]
	index int
	count int
}

func (it *Iterator[T]) Next() bool {
	if it.count >= it.deque.size {
		return false
	}
	it.index = (it.deque.head + it.count) % len(it.deque.data)
	it.count++
	return true
}

func (it *Iterator[T]) Value() T {
	return it.deque.data[it.index]
}

func (it *Iterator[T]) SetValue(value T) {
	it.deque.data[it.index] = value
}

func (it *Iterator[T]) Remove() {
	it.deque.Erase(it.index)
	it.index = -1
}
