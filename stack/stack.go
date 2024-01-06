package stack

type Node[T any] struct {
	Value T
	next  *Node[T]
}

type Stack[T any] struct {
	top  *Node[T]
	size int
}

type Iterator[T any] struct {
	current *Node[T]
}

func New[T any](values ...T) *Stack[T] {
	stack := &Stack[T]{}
	for _, v := range values {
		stack.Push(v)
	}
	return stack
}

func (s *Stack[T]) Push(value T) {
	newNode := &Node[T]{Value: value, next: s.top}
	s.top = newNode
	s.size++
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.top == nil {
		var zeroValue T
		return zeroValue, false
	}

	value := s.top.Value
	s.top = s.top.next
	s.size--

	return value, true
}

// Top returns the top element of the stack without removing it. Returns a zero value and false if the stack is empty.
func (s *Stack[T]) Top() (T, bool) {
	if s.size == 0 {
		var zeroValue T
		return zeroValue, false
	}

	return s.top.Value, true
}

func (s *Stack[T]) Size() int {
	return s.size
}

// Empty checks whether the stack is empty.
func (s *Stack[T]) Empty() bool {
	return s.size == 0
}

func (s *Stack[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{current: s.top}
}

func (it *Iterator[T]) Next() bool {
	if it.current == nil {
		return false
	}

	it.current = it.current.next

	return it.current != nil
}

func (it *Iterator[T]) Value() T {
	if it.current == nil {
		var zeroValue T
		return zeroValue
	}

	return it.current.Value
}
