package unordered_map

import "fmt"

func ExampleNew() {
	// Create a new unordered_map
	m := New[int, int](HashInt)

	// Add items to the unordered_map
	m.Set(1, 2)
	m.Set(2, 3)

	// Iterate over the unordered_map
	it := m.Iterator()
	for it.Next() {
		k, v := it.Value()
		fmt.Println(k, v)
	}

	// Output:
	// 1 2
	// 2 3
}
