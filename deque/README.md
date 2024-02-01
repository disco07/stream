# Go Deque

## Overview

The `deque` package in Go offers a generic implementation of a double-ended queue (Deque), drawing inspiration from the C++ standard library's Deque. It allows for efficient addition and removal of elements from both ends of the queue, providing a versatile data structure for a wide range of applications.

## Features

- **Generic Implementation**: Utilizes Go's type parameters to support elements of any type, enhancing flexibility and type safety.
- **Efficient Operations**: Designed for amortized constant time complexity for push and pop operations at both ends.
- **Dynamic Resizing**: Automatically adjusts its capacity to maintain operational efficiency as elements are added or removed.
- **Intuitive Usage**: Provides a clear and simple API for interacting with the Deque, including methods like `PushFront`, `PushBack`, `PopFront`, and `PopBack`.

## Installation

To use the `deque` package in your Go project, start by ensuring you have Go 1.18 or later installed. Then, include the package in your project by importing it:

```go
import "path/to/deque"
```

Replace "path/to/deque" with the actual path where the deque package is located.

Usage
Here's a simple example to demonstrate how to use the deque package:

```go
package main

import (
    "fmt"
    "path/to/deque"
)

func main() {
    // Initialize a new Deque with some integers
    dq := deque.New[int](1, 2, 3)

    // Add elements to the front and back
    dq.PushFront(0)
    dq.PushBack(4)

    // Remove and print elements from both ends
    fmt.Println(dq.PopFront()) // Outputs: 0
    fmt.Println(dq.PopBack())  // Outputs: 4
}
```