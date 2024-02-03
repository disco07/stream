// Package unordered_map implements a generic hash map data structure in Go, inspired by
// the C++ standard library's unordered_map. This Go implementation provides efficient
// key-value storage with fast lookup, addition, and deletion operations, optimized for
// performance through hash-based storage without maintaining element order.
//
// Drawing from the design principles of C++'s unordered_map, this package offers a
// Map type supporting generic key-value pairs. Keys must be of a comparable type, while
// values can be of any type, utilizing Go's type parameters for flexibility across
// various applications.
//
// Internally, the Map is realized as an array of keyValue structures, each encapsulating
// a key, value, and a boolean occupancy flag. This approach efficiently addresses hash
// collisions and conserves memory.
//
// Key Features:
//   - Customizable Hash Function: Users can specify their own hash function for key hashing,
//     accommodating diverse hashing needs and strategies.
//   - Dynamic Resizing: To ensure operational efficiency, the map's capacity dynamically
//     adjusts in response to changes in the number of stored elements.
//   - Generics: The use of generics enables the Map to handle any comparable key and arbitrary
//     value types, enhancing its utility in a wide array of use cases.
//
// Example Usage:
//
//	// A simple string hash function example.
//	hashFn := HashString
//
//	// Instantiate a Map with the custom hash function.
//	m := unordered_map.New[string, int](hashFn)
//
//	// Insert key-value pairs.
//	m.Set("apple", 5)
//	m.Set("banana", 10)
//	m.Set("cherry", 15)
//
//	// Retrieve the value for a key.
//	if value, ok := m.Get("banana"); ok {
//	    fmt.Println("Value:", value) // Expected Output: Value: 10
//	}
//
// This illustrative example demonstrates initializing a new unordered map with string keys
// and integer values, employing a custom hash function. It includes examples of adding
// key-value pairs to the map and retrieving a specific value.
//
// The design and functionality of this unordered_map package pay homage to the
// flexibility and efficiency of its C++ counterpart, adapted to Go's language
// idioms and type system.
package unordered_map

import (
	"github.com/segmentio/fasthash/fnv1a"
	"golang.org/x/exp/constraints"
	"math/rand"
	"reflect"
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

	hashFn func(K) uint64
}

type HashFn[K comparable] func(K) uint64

func New[K comparable, V any](hashFn HashFn[K]) *Map[K, V] {
	return &Map[K, V]{
		items:    make([]keyValue[K, V], initialCapacity),
		capacity: initialCapacity,
		hashFn:   hashFn,
	}
}

func (m *Map[K, V]) Set(key K, value V) {
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

func (m *Map[K, V]) Size() uint64 {
	return m.size
}

func (m *Map[K, V]) Empty() bool {
	return m.size == 0
}

func (m *Map[K, V]) Clear() {
	m.items = make([]keyValue[K, V], initialCapacity)
	m.capacity = initialCapacity
	m.size = 0
}

func (m *Map[K, V]) EraseIf(predicate func(K, V) bool) {
	for _, e := range m.items {
		if e.filled && predicate(e.key, e.value) {
			m.Delete(e.key)
		}
	}
}

func Contains[K comparable, V any](m *Map[K, V], key K) bool {
	_, ok := m.Get(key)

	return ok

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

type Hashable interface {
	constraints.Integer | constraints.Float | string
}

func Hash[T Hashable](t T) uint64 {
	switch v := any(t).(type) {
	case string:
		return hashString(v)
	}
	return hash(toUint64(t))
}

func hashString(s string) uint64 {
	return fnv1a.HashString64(s)
}

func toUint64(value interface{}) uint64 {
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint()
	case reflect.Float32, reflect.Float64:
		return uint64(val.Float())
	default:
		panic("unsupported type")
	}
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
