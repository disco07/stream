package hashmap

import (
	"cmp"
	"sort"
)

type KeyValuePair[K cmp.Ordered, V any] struct {
	Key   K
	Value V
}

type SortedMap[K cmp.Ordered, V any] struct {
	data []KeyValuePair[K, V]
}

func NewSortedMap[K cmp.Ordered, V any]() *SortedMap[K, V] {
	return &SortedMap[K, V]{}
}

func (m *SortedMap[K, V]) Insert(key K, value V) {
	i := sort.Search(len(m.data), func(i int) bool { return m.data[i].Key >= key })
	if i < len(m.data) && m.data[i].Key == key {
		m.data[i].Value = value
		return
	}
	m.data = append(m.data, KeyValuePair[K, V]{})
	copy(m.data[i+1:], m.data[i:])
	m.data[i] = KeyValuePair[K, V]{Key: key, Value: value}
}

func (m *SortedMap[K, V]) Get(key K) (V, bool) {
	i := sort.Search(len(m.data), func(i int) bool { return m.data[i].Key >= key })
	if i < len(m.data) && m.data[i].Key == key {
		return m.data[i].Value, true
	}
	var zeroValue V
	return zeroValue, false
}
