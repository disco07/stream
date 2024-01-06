package deque

type Node[T any] struct {
	Value T
	next  *Node[T]
	prev  *Node[T]
}

type Deque[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

type Iterator[T any] struct {
	current *Node[T]
}

func New[T any](values ...T) *Deque[T] {
	deque := &Deque[T]{}
	for _, v := range values {
		deque.PushBack(v)
	}
	return deque
}

func (d *Deque[T]) PushFront(value T) {
	newNode := &Node[T]{Value: value, next: d.head}
	if d.head != nil {
		d.head.prev = newNode
	} else {
		d.tail = newNode
	}
	d.head = newNode
	d.size++
}

func (d *Deque[T]) PushBack(value T) {
	newNode := &Node[T]{Value: value, prev: d.tail}
	if d.tail != nil {
		d.tail.next = newNode
	} else {
		d.head = newNode
	}
	d.tail = newNode
	d.size++
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{current: d.head}
}
