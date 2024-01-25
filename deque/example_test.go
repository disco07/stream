package deque

import "fmt"

func ExampleNew() {
	// Create a new deque
	d := New[int]()

	// Add items to the deque
	d.PushBack(1)
	d.PushBack(2)

	// Iterate over the deque
	it := d.Iterator()
	for it.Next() {
		fmt.Println(it.Value())
	}

	// Output:
	// 1
	// 2
}

func ExampleDeque_PushFront() {
	d := New[int]()

	d.PushFront(1)
	d.PushFront(2)

	it := d.Iterator()
	for it.Next() {
		fmt.Println(it.Value())
	}

	// Output:
	// 2
	// 1
}

func ExampleContains() {
	d := New[int]()
	d.PushBack(1)
	d.PushBack(2)

	fmt.Println(Contains(d, 1))
	fmt.Println(Contains(d, 3))

	// Output:
	// true
	// false
}

func ExampleSort() {
	d := New[int](4, 9, 4, 2, 5, 1, 3)

	Sort(d)

	it := d.Iterator()
	for it.Next() {
		fmt.Println(it.Value())
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 4
	// 5
	// 9
}
