// Package unordered_map provides a generic implementation of a hashmap, similar to
// std::unordered_map in C++. It's designed to offer a flexible and efficient
// way to manage key-value pairs with customizable hashing and equality operations.
//
// The Map structure at the core of this package handles key-value pairs using
// an array-based approach. The New function initializes the Map:
//
//	func New[K, V any](HashOps HashOps[K]) *Map[K, V] {
//	    return &Map[K, V]{
//	        items:    make([]keyValue[K, V], initialCapacity),
//	        capacity: initialCapacity,
//	        hashOps:  HashOps,
//	    }
//	}
//
// The HashOps parameter requires the user to provide hash and equality functions,
// enabling the use of the Map with any types for keys and values. This package
// is optimized for performance, with a focus on efficient memory usage. The map
// supports operations like Set, Get, and resizing, maintaining good performance
// characteristics even as it grows.
//
// Usage example:
//
//	func equalsInt(a, b int) bool {
//	    return a == b
//	}
//
//	func hashInt(a int) uint64 {
//	    return uint64(a)
//	}
//
//	func main() {
//	    hashOps := HashOps[int]{equals: equalsInt, hash: hashInt}
//	    m := NewMap[int, string](hashOps)
//	    m.Set(1, "one")
//	    m.Set(2, "two")
//
//	    if value, ok := m.Get(1); ok {
//	        fmt.Println("Value:", value)
//	    }
//	}
//
// unordered_map is ideal for applications requiring a hashmap with customized
// hashing and equality checks, and where performance and efficient memory usage are critical.
package unordered_map

import (
	"fmt"
	"hash/fnv"
	"math/rand"
)

var initialCapacity = uint64(rand.Intn(1024))

type keyValue[K comparable, V any] struct {
	key    K
	filled bool
	value  V
}

type Map[K comparable, V any] struct {
	items []keyValue[K, V]

	capacity uint64
	size     uint64
	readonly bool

	hashFn func(K) uint64
}

type HashOps[T any] struct {
	equals func(a, b T) bool
}

func New[K comparable, V any](hashFn func(K) uint64) *Map[K, V] {
	return &Map[K, V]{
		items:    make([]keyValue[K, V], initialCapacity),
		capacity: initialCapacity,
		hashFn:   hashFn,
	}
}

func (m *Map[K, V]) Set(key K, value V) {
	if m.readonly {
		panic("cannot modify readonly map")
	}

	if float64(m.size)/float64(m.capacity) > 0.75 {
		m.resize()
	}

	hash := m.hashFn(key)
	index := hash % m.capacity

	for {
		if !m.items[index].filled || m.items[index].key == key {
			m.items[index] = keyValue[K, V]{key, true, value}
			m.size++

			return
		}

		index = (index + 1) % m.capacity
	}
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	hash := m.hashFn(key)
	index := hash % m.capacity

	for {
		e := m.items[index]

		if e.filled {
			if e.key == key {
				return e.value, true
			}

			index = (index + 1) % m.capacity
		} else {
			var zero V

			return zero, false
		}
	}
}

func (m *Map[K, V]) Delete(key K) {
	if m.readonly {
		panic("cannot modify readonly map")
	}

	hash := m.hashFn(key)
	index := hash % m.capacity

	for {
		e := m.items[index]

		if e.filled {
			if e.key == key {
				m.items[index] = keyValue[K, V]{}
				m.size--

				return
			}

			index = (index + 1) % m.capacity
		} else {
			return
		}
	}
}

func (m *Map[K, V]) Iterator() *Iterator[K, V] {
	return &Iterator[K, V]{m: m, index: 0, count: 0}
}

func (m *Map[K, V]) resize() {
	newCapacity := m.capacity * 2
	newEntries := make([]keyValue[K, V], newCapacity)
	for _, e := range m.items {
		if e.filled {
			hash := m.hashFn(e.key)
			index := hash % newCapacity

			for {
				if !newEntries[index].filled {
					newEntries[index] = e
					break
				}

				index = (index + 1) % newCapacity
			}
		}
	}

	m.items = newEntries
	m.capacity = newCapacity
}

func hashInt(a int) uint64 {
	return uint64(a)
}

func hash(key any, cap uint64) uint64 {
	hasher := fnv.New64a()
	hasher.Write([]byte(fmt.Sprintf("%v", key)))

	return hasher.Sum64() % cap
}
