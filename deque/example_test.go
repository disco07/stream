package deque

import "fmt"

func ExampleNew() {
	// Create a new deque
	d := New[int]()

	// Add items to the deque
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	// Iterate over the deque
	it := d.Iterator()
	for it.Next() {
		fmt.Println(it.Value())
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleDeque_PushFront() {
	d := New[int]()

	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)

	it := d.Iterator()
	for it.Next() {
		fmt.Println(it.Value())
	}

	// Output:
	// 3
	// 2
	// 1
}
