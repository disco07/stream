package stack

type Iterator[T any] struct {
	stack *Stack[T]
	index int
}

func (it *Iterator[T]) Next() bool {
	it.index++
	return it.index < it.stack.size
}

// Value renvoie la valeur actuelle de l'itérateur
func (it *Iterator[T]) Value() T {
	return it.stack.data[it.index]
}

func (it *Iterator[T]) SetValue(value T) {
	if it.index < 0 || it.index >= it.stack.size {
		return
	}

	it.stack.data[it.index] = value
}

func (it *Iterator[T]) Remove() {
	if it.index < 0 || it.index >= it.stack.size {
		return
	}

	copy(it.stack.data[it.index:], it.stack.data[it.index+1:])
	it.stack.size--
	it.index--
}
