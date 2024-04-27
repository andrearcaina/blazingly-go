# Generics in Go

Generics are a way to write functions, data structures, and algorithms that can work with any type. They allow you to write code that is more flexible and reusable. Generics are a powerful feature that can help you write cleaner and more efficient code.

```go
package main

import "fmt"

// A generic function that takes a slice of any type and returns the sum of the elements in the slice.
func sum[T int | float64](s []T) T {
    var total T
    for _, v := range s {
        total += v
    }
    return total
}

func main() {
    // Create a slice of integers
    ints := []int{1, 2, 3, 4, 5}
    // Call the sum function with the slice of integers
    fmt.Println(sum(ints))

    // Create a slice of floats
    floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
    // Call the sum function with the slice of floats
    fmt.Println(sum(floats))
}

```