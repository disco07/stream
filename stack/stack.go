// Package stack provides a generic implementation of a stack data structure,
// supporting LIFO (Last In, First Out) semantics for element access.
//
// The Stack[T] structure is the primary type offered by this package, implemented
// using a slice to store elements. The New function initializes the Stack with an
// optional set of initial values:
//
//	func New[T any](values ...T) *Stack[T] {
//	    data := make([]T, len(values))
//	    copy(data, values)
//	    return &Stack[T]{data: data, size: len(values)}
//	}
//
// Stack supports typical stack operations such as Push, Pop, and Peek, making it
// suitable for various applications like expression evaluation, backtracking algorithms,
// and more.
//
// Example usage:
//
//	st := NewStack[int](1, 2)
//	st.Push(3)
//	fmt.Println(st.Pop()) // Outputs "3"
//
// This package is well-suited for use cases that require a simple and efficient
// last-in-first-out data management system.
package stack

type Stack[T any] struct {
	data []T
	size int
	tail int
	head int
}

// New creates a new stack with the given values.
func New[T any](items ...T) *Stack[T] {
	capacity := len(items) * 2
	if capacity < 4 {
		capacity = 4
	}

	stack := &Stack[T]{
		data: make([]T, capacity),
		size: len(items),
		tail: len(items),
		head: len(items),
	}

	copy(stack.data, items)
	return stack
}

// Push adds a value to the top of the stack.
func (s *Stack[T]) Push(item T) {
	if s.size == len(s.data) {
		s.resize()
	}
	s.data[s.tail] = item
	s.tail = s.tail + 1
	s.size++
}

// Pop removes the top value from the stack and returns it.
// Returns a zero value and false if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	var zeroValue T
	if s.size == 0 {
		return zeroValue, false
	}

	s.size--
	return s.data[s.size], true
}

// Top returns the top element of the stack without removing it.
// Returns a zero value and false if the stack is empty.
func (s *Stack[T]) Top() (T, bool) {
	var zeroValue T
	if s.size == 0 {
		return zeroValue, false
	}

	return s.data[s.size-1], true
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return s.size
}

// Empty returns true if the stack is empty.
func (s *Stack[T]) Empty() bool {
	return s.size == 0
}

func (s *Stack[T]) resize() {
	newData := make([]T, len(s.data)*2)

	copy(newData, s.data[s.head:])
	s.data = newData
	s.head = 0
	s.tail = s.size
}

// Iterator returns an iterator for the stack.
func (s *Stack[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{stack: s, index: s.size}
}
