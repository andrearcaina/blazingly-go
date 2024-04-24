# Pointers

## What is a pointer?

A pointer is a variable that stores the memory address of another variable. In other words, a pointer "points to" the location of a variable in memory.

## Pointers in Go

In Go, pointers are used to store the memory address of a value. The `&` operator is used to get the memory address of a variable, and the `*` operator is used to get the value that a pointer points to.

```go
package main

import "fmt"

func main() {
    x := 42
    p := &x // p is a pointer to x

    fmt.Println(*p) // prints the value that p points to (42)
}
```

If I were to print the value of p without dereferencing it, it would print the memory address of x.

```go
fmt.Println(p) // prints the memory address of x
```

This is unlike in C, where pointers are used to directly access memory locations and manipulate data. In Go, pointers are used to pass references to values, which can be useful when working with large data structures or when you want to modify the value of a variable in a function.

The same can be said for Rust. Rust dereferences pointers without the need for the `*` operator. This is because Rust has a concept called "deref coercion" that automatically dereferences pointers when needed. But in Go, you need to explicitly dereference pointers using the `*` operator.

## Pointers vs. Values

When you pass a value to a function in Go, a copy of the value is made and passed to the function. This means that any changes made to the value inside the function will not affect the original value.

```go
package main

import "fmt"

func double(x int) {
    x = x * 2
}

func main() {
    x := 42
    double(x)
    fmt.Println(x) // prints 42
}
```

If you want to modify the original value inside a function, you can pass a pointer to the value instead.

```go
package main

import "fmt"

func double(x *int) {
    *x = *x * 2
}

func main() {
    x := 42
    double(&x)
    fmt.Println(x) // prints 84
}
```

In this example, the `double` function takes a pointer to an integer as an argument and modifies the value that the pointer points to. By passing a pointer to the original value, we can modify the value inside the function and have the changes reflected in the original value.

## Slices and Maps vs Pointers

In Go, slices and maps are reference types, which means that when you pass a slice or map to a function, you are passing a reference to the underlying data structure, not a copy of the data structure.

This means that slices are already pointers, so you don't need to pass a pointer to a slice to modify the original slice inside a function.

Same can be said for maps.

But in functions, if you want to modify the slice or map itself (e.g., append elements to a slice or add elements to a map), you need to pass a pointer to the slice or map.

```go
package main

import "fmt"

func appendToSlice(s []int) {
    s = append(s, 4)
}

func appendToSlicePointer(s *[]int) {
    *s = append(*s, 4)
}

func main() {
    s := []int{1, 2, 3}
    appendToSlice(s)
    fmt.Println(s) // prints [1 2 3]

    appendToSlicePointer(&s)
	
    fmt.Println(s) // prints [1 2 3 4]
}
```