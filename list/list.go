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

func New[T any](values ...T) *List[T] {
	list := &List[T]{}
	for _, v := range values {
		list.PushBack(v)
	}

	return list
}

// PushBack ajoute un élément à la fin de la liste.
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

// PopFront supprime et renvoie l'élément du début de la liste.
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

// PopBack supprime et renvoie l'élément de fin de la liste.
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

// Size retourne le nombre d'éléments dans la liste.
func (l *List[T]) Size() int {
	return l.size
}

// Empty vérifie si la liste est vide.
func (l *List[T]) Empty() bool {
	return l.size == 0
}

// Clear efface tous les éléments de la liste.
func (l *List[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *List[T]) Front() (T, bool) {
	if l.head == nil {
		var zeroValue T
		return zeroValue, false
	}

	return l.head.Value, true
}

func (l *List[T]) Back() (T, bool) {
	if l.tail == nil {
		var zeroValue T
		return zeroValue, false
	}

	return l.tail.Value, true
}

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

func (l *List[T]) PrependRange(y *List[T]) {
	l.InsertRange(&Iterator[T]{current: l.head, size: l.size}, y)
}

func (l *List[T]) Resize(newSize int, zeroValue T) {
	for l.size < newSize {
		l.PushBack(zeroValue)
	}
	for l.size > newSize {
		l.PopBack()
	}
}

func (l *List[T]) Merge(y *List[T], compare func(a, b T) bool) {
	it := l.Begin()
	for itY := y.Begin(); itY.current != nil; y.Erase(itY) {
		for ; it.current != nil && compare(it.Value(), itY.Value()); it.Next() {
		}
		l.Insert(it, itY.Value())
	}
}

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

func (l *List[T]) Swap(y *List[T]) {
	l.head, y.head = y.head, l.head
	l.tail, y.tail = y.tail, l.tail
	l.size, y.size = y.size, l.size
}

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

func (l *List[T]) Sort(compare func(a, b T) bool) {
	if l.size < 2 {
		return
	}
	// Simple insertion sort, replace with more efficient sort for large lists
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

// Begin retourne un itérateur pointant sur le premier élément de la liste.
func (l *List[T]) Begin() *Iterator[T] {
	return &Iterator[T]{current: l.head, size: l.Size()}
}

// End retourne un itérateur pointant sur l'élément après le dernier élément de la liste.
func (l *List[T]) End() *Iterator[T] {
	return &Iterator[T]{current: nil, size: l.Size()}
}

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

	// Détacher la plage de nœuds de y.
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

	// Attacher la plage de nœuds à la liste l à la position pos.
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

	// Ajuster la taille des listes.
	count := 0
	for node := firstNode; node != lastNode.next; node = node.next {
		count++
	}
	y.size -= count
	l.size += count
}
