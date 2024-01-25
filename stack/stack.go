package stack

type Stack[T any] struct {
	data []T
	size int
}

func New[T any](values ...T) *Stack[T] {
	data := make([]T, len(values))
	copy(data, values)
	return &Stack[T]{data: data, size: len(values)}
}

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

func (s *Stack[T]) Pop() (T, bool) {
	var zeroValue T
	if s.size == 0 {
		return zeroValue, false
	}
	s.size--
	return s.data[s.size], true
}

// Top returns the top element of the stack without removing it. Returns a zero value and false if the stack is empty.
func (s *Stack[T]) Top() (T, bool) {
	var zeroValue T
	if s.size == 0 {
		return zeroValue, false
	}
	return s.data[s.size-1], true
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) Empty() bool {
	return s.size == 0
}

func (s *Stack[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{stack: s, index: -1}
}
