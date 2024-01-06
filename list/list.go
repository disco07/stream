package list

type Node[T any] struct {
	Value T
	next  *Node[T]
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

type Iterator[T any] struct {
	current *Node[T]
}

// New creates a new empty list.
func New[T any](values ...T) *List[T] {
	list := &List[T]{}
	for _, v := range values {
		list.Add(v)
	}
	return list
}

// Add appends a new element to the list.
func (l *List[T]) Add(value T) {
	newNode := &Node[T]{Value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

// Size returns the number of elements in the list.
func (l *List[T]) Size() int {
	return l.size
}

// Iterator returns an iterator for the list.
func (l *List[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{current: l.head}
}

// Next moves the iterator to the next element and returns true if there was a next element in the list.
func (it *Iterator[T]) Next() bool {
	if it.current == nil {
		return false
	}
	it.current = it.current.next
	return it.current != nil
}

// Value returns the current element of the iterator.
func (it *Iterator[T]) Value() T {
	if it.current == nil {
		var zeroValue T
		return zeroValue
	}
	return it.current.Value
}
