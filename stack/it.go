package stack

type Iterator[T any] struct {
	stack *Stack[T]
	index int
}

// Next returns true if there is a next value.
func (it *Iterator[T]) Next() bool {
	it.index++
	return it.index < it.stack.size
}

// Value returns the current value.
func (it *Iterator[T]) Value() T {
	return it.stack.data[it.index]
}

// SetValue sets the current value.
func (it *Iterator[T]) SetValue(value T) {
	if it.index < 0 || it.index >= it.stack.size {
		return
	}

	it.stack.data[it.index] = value
}

// Remove removes the current value.
func (it *Iterator[T]) Remove() {
	if it.index < 0 || it.index >= it.stack.size {
		return
	}

	copy(it.stack.data[it.index:], it.stack.data[it.index+1:])
	it.stack.size--
	it.index--
}
