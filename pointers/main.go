package main

import (
	"fmt"
)

func main() {
	var numb int32 = 10
	var num2 int32 = 1290315
	var pointer *int32 = &numb // or var pointer = &numb or pointer := &numb

	fmt.Println("Value of numb:", numb)
	fmt.Println("Address of numb:", &numb)
	fmt.Println("(Address of numb) value of pointer:", pointer) // doesn't dereference the pointer like Rust (just prints the address)
	fmt.Println("Actual value at the pointer's value (the value from the address):", *pointer)

	fmt.Println("----------------------------------")

	*pointer = 20 // dereferencing the pointer to change the value at the address

	fmt.Println("Value of numb:", numb)
	fmt.Println("Address of numb:", &numb)
	fmt.Println("(Address of numb) value of pointer:", pointer)
	fmt.Println("Actual value at the pointer's value (the value from the address):", *pointer)

	fmt.Println("----------------------------------")

	*pointer = *pointer / 10

	fmt.Println("Value of numb:", numb)
	fmt.Println("Address of numb:", &numb)
	fmt.Println("(Address of numb) value of pointer:", pointer)
	fmt.Println("Actual value at the pointer's value (the value from the address):", *pointer)

	fmt.Println("----------------------------------")

	pointer = &num2
	*pointer /= 5 // this changes the value at the address of num2 (since pointer is pointing to num2)

	num2 = 100 // now we change the value of num2 (which means *pointer technically changes as well)

	fmt.Println("Value of num2:", num2)
	fmt.Println("Address of num2:", &num2)
	fmt.Println("(Address of num2) value of pointer:", pointer)
	fmt.Println("Actual value at the pointer's value (the value from the address):", *pointer)

	fmt.Println("----------------------------------")

	var p *int32 // initialized to nil (zero value for pointers)
	var i int32

	p = new(int32) // new() function allocates memory (assigns an address to the pointer) for the type and returns a pointer to it

	*p = 10 // dereferencing the pointer to change the value at the address (default value of p after new() is 0)

	fmt.Println("Value of i:", i)
	fmt.Println("Address of i:", &i)
	fmt.Println("Value of p:", p)
	fmt.Println("Actual value at the pointer's value (the value from the address):", *p)

	fmt.Println("----------------------------------")

	x := 10

	double(x) // this doesn't change the value of x

	fmt.Println("Value of x:", x)

	doubleByPointer(&x) // this changes the value of x

	fmt.Println("Value of x:", x)

	// x is now 20 (from the previous doubleByPointer function call)
	x = doubleWithReturn(x)

	fmt.Println("Value of x:", x)

	fmt.Println("----------------------------------")

	var slice []int = []int{1, 2, 3, 4, 5}
	var sliceCopy = slice
	sliceCopy[2] = 10

	fmt.Println("Slice:", slice)
	fmt.Println("Slice Copy:", sliceCopy)

	fmt.Println("----------------------------------")

	appendToSlice(slice, 6) // this doesn't change the value of slice

	fmt.Println("Slice:", slice)
	fmt.Println("Slice Copy:", sliceCopy)

	appendToSliceByPointer(&slice, 6) // this changes the value of slice

	fmt.Println("Slice:", slice)
	fmt.Println("Slice Copy:", sliceCopy) // sliceCopy is not affected by the changes to slice (since slice is passed by value)
}

// this function doesn't change the value of x (it's passed by value)
// this means that the value of x is copied to the function's parameter
// and the function only changes the copy (doesn't change the original value)
// it doesn't return anything
func double(x int) {
	x = x * 2
}

// this function changes the value of x (it's passed by reference)
// this means that the address of x is passed to the function's parameter
// and the function changes the value at the address (changes the original value)
func doubleByPointer(x *int) {
	*x = *x * 2
}

// this is a pure function (doesn't have side effects)
// this also returns a copy of the input (doesn't change the input value)
func doubleWithReturn(x int) int {
	return x * 2
}

// this function doesn't change the value of x (it's passed by value)
// this means that the value of slice is copied to the function's parameter
func appendToSlice(slice []int, value int) {
	slice = append(slice, value)
}

// this function changes the value of slice (it's passed by reference)
// this means that the address of slice is passed to the function's parameter
// and the function changes the value at the address (changes the original value)
func appendToSliceByPointer(slice *[]int, value int) {
	*slice = append(*slice, value)
}
