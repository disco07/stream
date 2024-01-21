package deque

type Deque[T any] struct {
	// data est la slice interne qui stocke les éléments du Deque.
	// Son type est générique, permettant au Deque de stocker des éléments de n'importe quel type spécifié.
	data []T

	// head est l'indice de l'élément en début du Deque.
	// Il pointe vers l'emplacement du prochain élément à supprimer lors d'un PopFront,
	// ou vers l'endroit où insérer un nouvel élément lors d'un PushFront.
	head int

	// tail est l'indice juste après le dernier élément du Deque.
	// Il pointe vers l'emplacement où un nouvel élément doit être ajouté lors d'un PushBack.
	// Notez que tail est l'indice du premier emplacement vide après la fin actuelle du Deque.
	tail int

	// size représente le nombre d'éléments actuellement stockés dans le Deque.
	// C'est la différence entre les indices tail et head, ajustée pour les débordements.
	size int

	// hasHead est un booléen qui indique si le Deque a un élément en début.
	hasHead bool
}

// NewDeque creates a new deque with the given items.
func NewDeque[T any](items ...T) *Deque[T] {
	capacity := len(items) * 2
	if capacity < 4 {
		capacity = 4
	}
	dq := &Deque[T]{
		data: make([]T, capacity),
		head: 0,
		tail: len(items),
		size: len(items),
	}
	copy(dq.data, items)
	return dq
}

func (d *Deque[T]) PushFront(item T) {
	d.hasHead = true
	if d.size == len(d.data) {
		d.resize()
	}
	d.head = (d.head - 1 + len(d.data)) % len(d.data)
	d.data[d.head] = item
	d.size++
}

func (d *Deque[T]) PushBack(item T) {
	if d.size == len(d.data) {
		d.resize()
	}
	d.data[d.tail] = item
	d.tail = d.tail + 1
	d.size++
}

func (d *Deque[T]) PopFront() (item T, ok bool) {
	if d.size == 0 {
		return
	}
	item = d.data[d.head]
	d.head = (d.head + 1) % len(d.data)
	d.size--
	ok = true
	return
}

func (d *Deque[T]) PopBack() (item T, ok bool) {
	if d.size == 0 {
		return
	}
	d.tail = (d.tail - 1 + len(d.data)) % len(d.data)
	item = d.data[d.tail]
	d.size--
	ok = true
	return
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) Empty() bool {
	return d.size == 0
}

func (d *Deque[T]) Front() (item T, ok bool) {
	if d.size == 0 {
		return
	}
	item = d.data[d.head]
	ok = true
	return
}

func (d *Deque[T]) Back() (item T, ok bool) {
	if d.size == 0 {
		return
	}
	item = d.data[(d.tail-1+len(d.data))%len(d.data)]
	ok = true
	return
}

func (d *Deque[T]) Clear() {
	d.head = 0
	d.tail = 0
	d.size = 0
	d.hasHead = false
}

func (d *Deque[T]) At(i int) (item T, ok bool) {
	if i < 0 || i >= d.size {
		return
	}
	item = d.data[(d.head+i)%len(d.data)]
	ok = true
	return
}

func (d *Deque[T]) Set(i int, item T) (ok bool) {
	if i < 0 || i >= d.size {
		return
	}
	d.data[(d.head+i)%len(d.data)] = item
	ok = true
	return
}

func (d *Deque[T]) Insert(i int, item T) (ok bool) {
	if i < 0 || i > d.size {
		return
	}
	if d.size == len(d.data) {
		d.resize()
	}
	if i < d.size/2 {
		d.head = (d.head - 1 + len(d.data)) % len(d.data)
		for j := 0; j < i; j++ {
			d.data[(d.head+j)%len(d.data)] = d.data[(d.head+j+1)%len(d.data)]
		}
	} else {
		for j := d.size; j > i; j-- {
			d.data[(d.head+j)%len(d.data)] = d.data[(d.head+j-1)%len(d.data)]
		}
	}
	d.data[(d.head+i)%len(d.data)] = item
	d.tail = (d.tail + 1) % len(d.data)
	d.size++
	return true
}

func (d *Deque[T]) InsertRange(i int, items []T) (ok bool) {
	if i < 0 || i > d.size {
		return
	}
	if d.size+len(items) > len(d.data) {
		d.resize()
	}
	if i < d.size/2 {
		d.head = (d.head - len(items) + len(d.data)) % len(d.data)
		for j := 0; j < i; j++ {
			d.data[(d.head+j)%len(d.data)] = d.data[(d.head+j+len(items))%len(d.data)]
		}
	} else {
		for j := d.size; j > i; j-- {
			d.data[(d.head+j+len(items))%len(d.data)] = d.data[(d.head+j)%len(d.data)]
		}
	}
	copy(d.data[(d.head+i)%len(d.data):], items)
	d.tail = (d.tail + len(items)) % len(d.data)
	d.size += len(items)
	return true
}

func (d *Deque[T]) Erase(i int) (ok bool) {
	if i < 0 || i >= d.size {
		return
	}
	if i < d.size/2 {
		for j := i; j > 0; j-- {
			d.data[(d.head+j)%len(d.data)] = d.data[(d.head+j-1)%len(d.data)]
		}
		d.head = (d.head + 1) % len(d.data)
	} else {
		for j := i; j < d.size-1; j++ {
			d.data[(d.head+j)%len(d.data)] = d.data[(d.head+j+1)%len(d.data)]
		}
		d.tail = (d.tail - 1 + len(d.data)) % len(d.data)
	}
	d.size--
	return true
}

func (d *Deque[T]) EraseRange(i, j int) (ok bool) {
	if i < 0 || i > j || j > d.size {
		return
	}
	if i < d.size/2 {
		for k := i; k > 0; k-- {
			d.data[(d.head+k)%len(d.data)] = d.data[(d.head+k-1)%len(d.data)]
		}
		d.head = (d.head + j - i) % len(d.data)
	} else {
		for k := j; k < d.size; k++ {
			d.data[(d.head+k-j+i)%len(d.data)] = d.data[(d.head+k)%len(d.data)]
		}
		d.tail = (d.tail - (j - i) + len(d.data)) % len(d.data)
	}
	d.size -= j - i
	return true
}

func (d *Deque[T]) AppendRange(items ...T) {
	if d.size+len(items) > len(d.data) {
		d.resize()
	}
	copy(d.data[d.tail:], items)
	d.tail = (d.tail + len(items)) % len(d.data)
	d.size += len(items)
}

func (d *Deque[T]) PrependRange(items ...T) {
	if d.size+len(items) > len(d.data) {
		d.resize()
	}
	d.head = (d.head - len(items) + len(d.data)) % len(d.data)
	copy(d.data[d.head:], items)
	d.size += len(items)
}

func (d *Deque[T]) Swap(i, j int) {
	d.data[(d.head+i)%len(d.data)], d.data[(d.head+j)%len(d.data)] = d.data[(d.head+j)%len(d.data)], d.data[(d.head+i)%len(d.data)]
}

func Contains[T comparable](d *Deque[T], item T) bool {
	for it := d.Iterator(); it.Next(); {
		if it.Value() == item {
			return true
		}
	}
	return false
}

func (d *Deque[T]) resize() {
	newData := make([]T, len(d.data)*2)

	copy(newData, d.data[d.head:])
	if d.hasHead {
		copy(newData[len(d.data[d.head:]):], d.data[:d.tail])
	}

	d.data = newData
	d.head = 0
	d.tail = d.size
}

type Iterator[T any] struct {
	deque *Deque[T]
	index int
	count int
}

func (it *Iterator[T]) Next() bool {
	if it.count >= it.deque.size {
		return false
	}
	it.index = (it.deque.head + it.count) % len(it.deque.data)
	it.count++
	return true
}

func (it *Iterator[T]) Value() T {
	return it.deque.data[it.index]
}

func (d *Deque[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{deque: d, index: -1, count: 0}
}
