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
	fmt.Println("(Address of numb) value of pointer:", pointer) // doesn't dereference the pointer like Rust
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
}
