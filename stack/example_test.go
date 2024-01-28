package stack

import (
	"fmt"
)

func ExampleNew() {
	stack := New[int](1, 2, 3)

	stack.Push(4)

	it := stack.Iterator()
	for it.Next() {
		fmt.Println(it.Value())
	}

	// Output:
	// 4
	// 3
	// 2
	// 1
}
