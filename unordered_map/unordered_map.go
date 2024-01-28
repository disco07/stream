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
	"github.com/segmentio/fasthash/fnv1a"
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

func HashUint64(u uint64) uint64 {
	return hash(u)
}
func HashUint32(u uint32) uint64 {
	return hash(uint64(u))
}
func HashUint16(u uint16) uint64 {
	return hash(uint64(u))
}
func HashUint8(u uint8) uint64 {
	return hash(uint64(u))
}
func HashInt64(i int64) uint64 {
	return hash(uint64(i))
}
func HashInt32(i int32) uint64 {
	return hash(uint64(i))
}
func HashInt16(i int16) uint64 {
	return hash(uint64(i))
}
func HashInt8(i int8) uint64 {
	return hash(uint64(i))
}
func HashInt(i int) uint64 {
	return hash(uint64(i))
}
func HashUint(i uint) uint64 {
	return hash(uint64(i))
}
func HashString(s string) uint64 {
	return fnv1a.HashString64(s)
}
func HashBytes(b []byte) uint64 {
	return fnv1a.HashBytes64(b)
}

func hash(u uint64) uint64 {
	u ^= u >> 33
	u *= 0xff51afd7ed558ccd
	u ^= u >> 33
	u *= 0xc4ceb9fe1a85ec53
	u ^= u >> 33
	u += u << 3
	u ^= u >> 7
	u *= 0x9e3779b97f4a7c15
	u ^= u >> 33

	return u
}
