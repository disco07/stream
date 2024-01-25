// Package stack implements a generic stack.
//
// This package requires Go 1.18 or later due to the use of generics. The Stack package
// offers a simple yet flexible implementation of a stack data structure, allowing for
// LIFO (Last In, First Out) operations. This implementation provides methods for
// pushing to and popping from the stack, as well as viewing the top element and checking
// if the stack is empty.
//
// The Stack is particularly useful in scenarios where you need to reverse the order of
// elements, keep track of previous states, or need a simple and efficient way to manage
// data with LIFO access patterns.
package stack

type Stack[T any] struct {
	data []T
	size int
}

// New creates a new stack with the given values.
func New[T any](values ...T) *Stack[T] {
	data := make([]T, len(values))
	copy(data, values)
	return &Stack[T]{data: data, size: len(values)}
}

// Push adds a value to the top of the stack.
func (s *Stack[T]) Push(value T) {
	if s.size == len(s.data) {
		// Double la capacité de la slice si nécessaire
		newCapacity := 2*s.size + 1
		if newCapacity < 4 {
			newCapacity = 4
		}

		newdata := make([]T, s.size, newCapacity)
		copy(newdata, s.data)
		s.data = newdata
	}
	s.data = s.data[:s.size+1]
	s.data[s.size] = value
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

// Iterator returns an iterator for the stack.
func (s *Stack[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{stack: s, index: -1}
}
