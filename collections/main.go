package main

import (
	"fmt"
)

func main() {
	var array [5]int // array of a fixed size of (5) integers (all initialized as 0)

	array[0] = 10

	fmt.Println(array)
	fmt.Printf("Memory Location of array: %v\n", &array)
	fmt.Printf("Length of array: %v\n", len(array))

	var array2 = [5]int{1, 2, 3, 4, 5}

	fmt.Println(array2)

	// array of arrays
	var array3 = [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(array3)
	fmt.Println(array3[0][1]) // 2 (accessing elements is like python)

	someArr := [...]int{-5, 123, 10, 76, 99} // array of unknown size (compiler will determine size)

	fmt.Println(someArr)

	// slices are arrays with dynamic sizes (and more functionality)
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice)

	fmt.Printf("Length of slice: %v, with capacity: %v\n", len(slice), cap(slice))
	slice = append(slice, 6) // append to slice (returns a new slice with the appended element)

	// capacity doubles when exceeded (to avoid frequent resizing), but length increases by 1
	// capacity = 2 * (previous capacity) when exceeded
	// can't actually access the capacity of a slice, but can use the cap() function to get it (like len())
	fmt.Printf("Length of slice: %v, with capacity: %v\n", len(slice), cap(slice))
	fmt.Println(slice)

	var slice2 = make([]int, 5) // make a slice with a length of 5 and using make() function
	var slice3 = new([5]int)[:] // make a slice with a length of 5 using new() function and [:] to convert to slice

	/*
		- make() function initializes the slice with the specified length and capacity
		- new() function initializes the slice with the specified length, but the capacity is not specified
		- new() function returns a pointer to the slice, basically creating an array and returning a pointer to it
		- [:] is used to convert the array to a slice (since new() returns a pointer to the array)
	*/

	fmt.Println(slice2)
	fmt.Println(slice3)

	slice = append(slice, slice2...) // append multiple slices to a slice
	slice = append(slice, slice3...)

	fmt.Println(slice)

	var slice4 []int32 = make([]int32, 5, 11) // make a slice with a length of 5 and a capacity of 8

	fmt.Printf("Length of slice: %v, with capacity: %v\n", len(slice4), cap(slice4))

	// slicing a slice

	slice5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(slice5[1:4]) // [2 3 4] (slicing a slice)
	fmt.Println(slice5[:4])  // [1 2 3 4] (slicing from the beginning)
	fmt.Println(slice5[4:])  // [5 6 7 8 9] (slicing to the end)
	fmt.Println(slice5[:])   // [1 2 3 4 5 6 7 8 9] (slicing the whole slice)

	// copying a slice
	slice6 := make([]int, 5)

	copy(slice6, slice5) // copy the elements of slice5 to slice6

	fmt.Println(slice6)

	// deleting elements from a slice
	slice7 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice7 = append(slice7[:2], slice7[3:]...) // delete the element at index 2

	fmt.Println(slice7)

	var myMap map[string]uint8 = make(map[string]uint8) // [string] == key, uint8 == value
	fmt.Printf("%v\n", myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Printf("%v\n", myMap2)
	fmt.Printf("%v, %v\n", myMap2["Adam"], myMap2["Sarah"])

	// maps always return true
	fmt.Printf("%v\n", myMap2["John"]) // 0 (default value for uint8)

	// So, do this instead
	var age, ok = myMap["John"]

	if ok {
		fmt.Printf("Age of John: %v\n", age)
	} else {
		fmt.Println("John not found")
	}

	delete(myMap2, "Adam") // delete an element from a map
	myMap2["John"] = 34    // add an element to a map

	fmt.Printf("%v\n", myMap2)

	// iterating over a slice
	for i, v := range slice7 {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// iterating over a slice without using the index
	for _, v := range slice7 {
		fmt.Printf("Value: %v\n", v)
	}

	// iterating over a slice without using the value
	for i := range slice7 {
		fmt.Printf("Index: %v\n", i)
	}

	// iterating over a map
	for k, v := range myMap2 {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}

	// iterating over a map without using the key
	for _, v := range myMap2 {
		fmt.Printf("Value: %v\n", v)
	}

	// iterating over a map without using the value
	for k := range myMap2 {
		fmt.Printf("Key: %v\n", k)
	}
}
