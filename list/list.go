// Package list implements a generic doubly linked list.
//
// This package requires Go 1.18 or later, as it makes use of generics introduced in Go 1.18.
// The List package provides a variety of methods to perform operations such as
// insertion, deletion, and traversal on a doubly linked list. This implementation
// allows for efficient insertion and removal from both the front and back of the list.
// The package includes an iterator to facilitate list traversal.
//
// Typical use cases of the List package include implementing queues, stacks, and other
// data structures that benefit from quick insertions and deletions at both ends.
package list

type Node[T any] struct {
	Value T
	// Remarque : Les champs prev et next restent des pointeurs car ils sont essentiels
	// pour la structure de la liste doublement chaînée.
	prev *Node[T]
	next *Node[T]
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// New create a new list with the given values.
func New[T any](values ...T) *List[T] {
	list := &List[T]{}
	for _, v := range values {
		list.PushBack(v)
	}

	return list
}

// PushBack adds a value to the end of the list.
func (l *List[T]) PushBack(value T) {
	newNode := &Node[T]{Value: value}
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}

	l.size++
}

// PopFront removes the first value of the list.
func (l *List[T]) PopFront() (T, bool) {
	if l.head == nil {
		var zeroValue T
		return zeroValue, false
	}
	value := l.head.Value
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	l.size--
	return value, true
}

// PopBack removes the last value of the list.
func (l *List[T]) PopBack() (T, bool) {
	if l.tail == nil {
		var zeroValue T
		return zeroValue, false
	}
	value := l.tail.Value
	l.tail = l.tail.prev
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	l.size--
	return value, true
}

// Size returns the number of elements in the list.
func (l *List[T]) Size() int {
	return l.size
}

// Empty returns true if the list is empty.
func (l *List[T]) Empty() bool {
	return l.size == 0
}

// Clear removes all elements from the list.
func (l *List[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Front returns the first element of the list.
func (l *List[T]) Front() (T, bool) {
	if l.head == nil {
		var zeroValue T
		return zeroValue, false
	}

	return l.head.Value, true
}

// Back returns the last element of the list.
func (l *List[T]) Back() (T, bool) {
	if l.tail == nil {
		var zeroValue T
		return zeroValue, false
	}

	return l.tail.Value, true
}

// Insert inserts a value at the given index.
func (l *List[T]) Insert(it *Iterator[T], value T) {
	newNode := &Node[T]{Value: value}
	if it.current != nil {
		newNode.next = it.current
		newNode.prev = it.current.prev
		if it.current.prev != nil {
			it.current.prev.next = newNode
		} else {
			l.head = newNode
		}
		it.current.prev = newNode
	} else {
		l.PushBack(value)
	}
	l.size++
}

// Erase removes the element at the given index.
func (l *List[T]) Erase(it *Iterator[T]) {
	if it.current == nil {
		return
	}
	if it.current.prev != nil {
		it.current.prev.next = it.current.next
	} else {
		l.head = it.current.next
	}
	if it.current.next != nil {
		it.current.next.prev = it.current.prev
	} else {
		l.tail = it.current.prev
	}
	l.size--
}

// Reverse reverses the list.
func (l *List[T]) Reverse() {
	var prev, next *Node[T]
	current := l.head
	l.tail = l.head
	for current != nil {
		next = current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}
	l.head = prev
}

// InsertRange inserts a list at the given index.
func (l *List[T]) InsertRange(it *Iterator[T], y *List[T]) {
	if y.head == nil {
		return
	}

	if it.current == nil {
		l.tail.next = y.head
		y.head.prev = l.tail
		l.tail = y.tail
	} else {
		if it.current.prev != nil {
			it.current.prev.next = y.head
			y.head.prev = it.current.prev
		} else {
			l.head = y.head
		}
		y.tail.next = it.current
		it.current.prev = y.tail
	}

	l.size += y.size
	y.head, y.tail, y.size = nil, nil, 0
}

// PrependRange prepends a list to the list.
func (l *List[T]) PrependRange(y *List[T]) {
	l.InsertRange(&Iterator[T]{current: l.head, size: l.size}, y)
}

// Resize resizes the list to the given size.
func (l *List[T]) Resize(newSize int, zeroValue T) {
	for l.size < newSize {
		l.PushBack(zeroValue)
	}
	for l.size > newSize {
		l.PopBack()
	}
}

// Merge merges two lists together.
func (l *List[T]) Merge(y *List[T], compare func(a, b T) bool) {
	it := l.Begin()
	for itY := y.Begin(); itY.current != nil; y.Erase(itY) {
		for ; it.current != nil && compare(it.Value(), itY.Value()); it.Next() {
		}
		l.Insert(it, itY.Value())
	}
}

// Unique removes all duplicate elements from the list.
func (l *List[T]) Unique(equal func(a, b T) bool) {
	it := l.Begin()
	if it.current == nil {
		return
	}
	for next := it.current.next; next != nil; next = it.current.next {
		if equal(it.Value(), next.Value) {
			l.Erase(&Iterator[T]{current: next, size: l.Size()})
		} else {
			it.current = next
		}
	}
}

// Swap swaps the contents of two lists.
func (l *List[T]) Swap(y *List[T]) {
	l.head, y.head = y.head, l.head
	l.tail, y.tail = y.tail, l.tail
	l.size, y.size = y.size, l.size
}

// RemoveIf removes all elements that satisfy the given predicate.
func (l *List[T]) RemoveIf(predicate func(T) bool) {
	for it := l.Begin(); it.current != nil; {
		if predicate(it.Value()) {
			next := it.current.next
			l.Erase(it)
			it.current = next
		} else {
			it.Next()
		}
	}
}

// Sort sorts the list using the given comparator.
func (l *List[T]) Sort(compare func(a, b T) bool) {
	if l.size < 2 {
		return
	}

	for it1 := l.Begin(); it1.Next(); {
		val := it1.Value()
		it2 := &Iterator[T]{current: it1.current.prev, size: l.Size()}
		for ; it2.current != nil && compare(val, it2.Value()); it2.current = it2.current.prev {
		}
		if it1.current != it2.current.next {
			l.Erase(it1)
			l.Insert(it2, val)
		}
	}
}

// Begin returns an iterator to the first element of the list.
func (l *List[T]) Begin() *Iterator[T] {
	return &Iterator[T]{current: l.head, size: l.Size()}
}

// End returns an iterator to the last element of the list.
func (l *List[T]) End() *Iterator[T] {
	return &Iterator[T]{current: nil, size: l.Size()}
}

// Splice removes the elements in the range [first, last) from y and inserts them into the list at pos.
func (l *List[T]) Splice(pos *Iterator[T], y *List[T], iterators ...*Iterator[T]) {
	if y.Empty() {
		return
	}

	var firstNode, lastNode *Node[T]

	switch len(iterators) {
	case 0:
		firstNode, lastNode = y.head, y.tail
	case 1:
		firstNode, lastNode = iterators[0].current, y.tail
	case 2:
		firstNode, lastNode = iterators[0].current, iterators[1].current.prev
	}

	if firstNode.prev != nil {
		firstNode.prev.next = lastNode.next
	} else {
		y.head = lastNode.next
	}
	if lastNode.next != nil {
		lastNode.next.prev = firstNode.prev
	} else {
		y.tail = firstNode.prev
	}

	if pos.current != nil {
		if pos.current.prev != nil {
			pos.current.prev.next = firstNode
			firstNode.prev = pos.current.prev
		} else {
			l.head = firstNode
		}
		lastNode.next = pos.current
		pos.current.prev = lastNode
	} else {
		if l.tail != nil {
			l.tail.next = firstNode
			firstNode.prev = l.tail
			l.tail = lastNode
		} else {
			l.head = firstNode
			l.tail = lastNode
		}
	}

	count := 0
	for node := firstNode; node != lastNode.next; node = node.next {
		count++
	}
	y.size -= count
	l.size += count
}
