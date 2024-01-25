package list

// Iterator est un itérateur pour List.
type Iterator[T any] struct {
	current *Node[T]
	size    int
	list    *List[T]
}

func (l *List[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{current: l.head, size: l.Size(), list: l}
}

// Next avance l'itérateur et retourne true s'il y a un élément suivant.
func (it *Iterator[T]) Next() bool {
	if it.current == nil {
		return false
	}

	if it.size > 1 {
		it.current = it.current.next
	}

	it.size--

	return it.size >= 0
}

// Value retourne la valeur courante de l'itérateur.
func (it *Iterator[T]) Value() T {
	if it.size < 0 {
		var zeroValue T
		return zeroValue
	}

	if it.size == 0 {
		return it.current.Value
	}

	return it.current.prev.Value
}

// SetValue modifie la valeur courante de l'itérateur.
func (it *Iterator[T]) SetValue(value T) {
	if it.size < 0 {
		return
	}

	if it.size == 0 {
		it.current.Value = value
		return
	}

	it.current.prev.Value = value
}

// Remove supprime l'élément courant de la liste.
func (it *Iterator[T]) Remove() {
	if it.size < 0 {
		return
	}

	if it.size == 0 {
		it.list.PopFront()
		return
	}

	if it.size == 1 {
		it.list.PopBack()
		return
	}

	it.current.prev.next = it.current.next
	it.current.next.prev = it.current.prev
	it.list.size--
}
