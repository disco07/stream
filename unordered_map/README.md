# Go Unordered Map

## Overview

The `unordered_map` library in Go is inspired by the C++ standard library's `unordered_map`. It provides an efficient hash map implementation that supports fast lookups, insertions, and deletions. The library utilizes Go's generics, introduced in Go 1.18, allowing for flexible key-value pair types.

## Features

- **Generic Implementation**: Utilizes Go's type parameters to allow keys and values of any comparable and any type respectively.
- **Customizable Hash Function**: Users can define their own hash function for keys, providing flexibility for various hashing needs.
- **Dynamic Resizing**: The capacity of the map dynamically adjusts to maintain efficient performance as elements are added or removed.
- **Efficient Collision Handling**: Implements an efficient method for handling hash collisions, ensuring consistent performance.

## Usage

To use the `unordered_map` library, first ensure you have Go 1.18 or later installed. Then, you can include the library in your project by importing it:

```go
package main

import (
    "fmt"
    "stream/unordered_map"
)

func main() {
    // Define a simple hash function for strings
    hashFn := unordered_map.HashString

    // Create a new Map instance
    m := unordered_map.New[string, int](hashFn)

    // Add key-value pairs
    m.Set("apple", 5)
    m.Set("banana", 10)

    // Retrieve a value
    if value, ok := m.Get("banana"); ok {
        fmt.Println("Value:", value) // Output: Value: 10
    }
}
```

