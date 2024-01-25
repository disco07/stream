// Package deque implements a generic deque (double-ended queue) data structure.
// This implementation allows for efficient addition and removal of elements from both
// the front and the back of the queue. It is implemented as a dynamically resizing circular buffer,
// providing amortized constant time complexity for the push and pop operations.
//
// The Deque is generic and can hold elements of any type, thanks to Go's type parameter feature introduced in Go 1.18.
// This feature makes the Deque flexible and versatile for various use cases.
//
// Key operations include PushFront and PushBack for adding elements, PopFront and PopBack for removing elements,
// and an At method for indexed access to elements. Additionally, the package provides a Sort function
// to sort the Deque in place using the QuickSort algorithm.
//
// The Deque is particularly useful in scenarios where both ends of the queue need to be accessed or modified,
// such as in certain algorithms or data processing tasks.
package deque

import (
	"cmp"
	"math/rand"
	"time"
)

// Deque is a double-ended queue.
type Deque[T any] struct {
	data    []T
	head    int
	tail    int
	size    int
	hasHead bool
}

// New creates a new deque with the given items.
func New[T any](items ...T) *Deque[T] {
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

// PushFront adds an item to the front of the deque.
func (d *Deque[T]) PushFront(item T) {
	d.hasHead = true
	if d.size == len(d.data) {
		d.resize()
	}
	d.head = (d.head - 1 + len(d.data)) % len(d.data)
	d.data[d.head] = item
	d.size++
}

// PushBack adds an item to the back of the deque.
func (d *Deque[T]) PushBack(item T) {
	if d.size == len(d.data) {
		d.resize()
	}
	d.data[d.tail] = item
	d.tail = d.tail + 1
	d.size++
}

// PopFront removes and returns an item from the front of the deque.
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

// PopBack removes and returns an item from the back of the deque.
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

// Size returns the number of items in the deque.
func (d *Deque[T]) Size() int {
	return d.size
}

// Empty returns true only if the deque is empty.
func (d *Deque[T]) Empty() bool {
	return d.size == 0
}

// Front returns the item at the front of the deque without removing it.
func (d *Deque[T]) Front() (item T, ok bool) {
	if d.size == 0 {
		return
	}
	item = d.data[d.head]
	ok = true
	return
}

// Back returns the item at the back of the deque without removing it.
func (d *Deque[T]) Back() (item T, ok bool) {
	if d.size == 0 {
		return
	}
	item = d.data[(d.tail-1+len(d.data))%len(d.data)]
	ok = true
	return
}

// Clear removes all items from the deque.
func (d *Deque[T]) Clear() {
	d.head = 0
	d.tail = 0
	d.size = 0
	d.hasHead = false
}

// At returns the item at the given index without removing it.
func (d *Deque[T]) At(i int) (item T, ok bool) {
	if i < 0 || i >= d.size {
		return
	}
	item = d.data[(d.head+i)%len(d.data)]
	ok = true
	return
}

// Set sets the item at the given index.
func (d *Deque[T]) Set(i int, item T) (ok bool) {
	if i < 0 || i >= d.size {
		return
	}
	d.data[(d.head+i)%len(d.data)] = item
	ok = true
	return
}

// Insert inserts an item at the given index.
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

// InsertRange inserts a slice of items at the given index.
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

// Erase removes the item at the given index.
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

// EraseRange removes a slice of items starting at the given index.
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

// AppendRange appends a slice of items to the back of the deque.
func (d *Deque[T]) AppendRange(items ...T) {
	if d.size+len(items) > len(d.data) {
		d.resize()
	}
	copy(d.data[d.tail:], items)
	d.tail = (d.tail + len(items)) % len(d.data)
	d.size += len(items)
}

// PrependRange prepends a slice of items to the front of the deque.
func (d *Deque[T]) PrependRange(items ...T) {
	if d.size+len(items) > len(d.data) {
		d.resize()
	}
	d.head = (d.head - len(items) + len(d.data)) % len(d.data)
	copy(d.data[d.head:], items)
	d.size += len(items)
}

// Swap swaps the items at the given indices.
func (d *Deque[T]) Swap(i, j int) {
	d.data[(d.head+i)%len(d.data)], d.data[(d.head+j)%len(d.data)] = d.data[(d.head+j)%len(d.data)], d.data[(d.head+i)%len(d.data)]
}

// EraseIf removes all items that satisfy the given predicate.
func (d *Deque[T]) EraseIf(predicate func(T) bool) {
	for it := d.Iterator(); it.Next(); {
		if predicate(it.Value()) {
			it.Remove()
		}
	}
}

// Contains returns true if the deque contains the given item.
func Contains[T comparable](d *Deque[T], item T) bool {
	for it := d.Iterator(); it.Next(); {
		if it.Value() == item {
			return true
		}
	}
	return false
}

// Sort sorts the deque in-place.
func Sort[T cmp.Ordered](d *Deque[T]) {
	if d.size > 1 {
		quickSort(d, 0, d.size-1)
	}
}

func quickSort[T cmp.Ordered](d *Deque[T], low, high int) {
	if low < high {
		pi := partition(d, low, high)
		quickSort(d, low, pi-1)
		quickSort(d, pi+1, high)
	}
}

func partition[T cmp.Ordered](d *Deque[T], low, high int) int {
	rand.Seed(time.Now().UnixNano())
	pivotIndex := rand.Intn(high-low+1) + low
	pivotValue, _ := d.At(pivotIndex)
	d.Swap(pivotIndex, high)
	i := low
	for j := low; j < high; j++ {
		if val, _ := d.At(j); val < pivotValue {
			d.Swap(i, j)
			i++
		}
	}
	d.Swap(i, high)
	return i
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

// Iterator is an iterator over a deque's items.
func (d *Deque[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{deque: d, index: -1, count: 0}
}
