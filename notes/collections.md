# Collections

Collections are data structures that store multiple items. Go has three built-in collections: arrays, slices, and maps.

## Arrays
Arrays are fixed-size collections of items of the same type. 

Arrays are declared using the following syntax: `var array [size]type` 

Where:
- array is the name of the array
- size is the number of elements in the array
- type is the type of the elements in the array

### Additional Information
- Arrays are zero-indexed, meaning that the first element is at index 0.
- Arrays are initialized with the zero value of the type of the elements.
- Arrays are passed by value, meaning that a copy of the array is passed to functions.
- Arrays are not commonly used in Go, slices are more commonly used.
  - This is because slices are more flexible and easier to work with, and slices are basically dynamic arrays.

## Slices

Slices are a flexible and dynamic collection of items of the same type. \
Slices are built on top of arrays.

Slices are declared using the following syntax: `var slice []type`

Where:
- slice is the name of the slice
- type is the type of the elements in the slice
- The size of the slice is not specified in the declaration
- Slices are created using the built-in `make` function

### Additional Information
- Slices are zero-indexed, meaning that the first element is at index 0.
- Slices are passed by reference, meaning that a reference to the underlying array is passed to functions.

## Maps

Maps are collections of key-value pairs. \
Maps are declared using the following syntax: `var mapName map[keyType]valueType`

Where:
- mapName is the name of the map
- keyType is the type of the keys in the map
- valueType is the type of the values in the map
- The size of the map is not specified in the declaration
- Maps are created using the built-in `make` function

### Additional Information
- Maps are unordered collections of key-value pairs.
- Maps are passed by reference, meaning that a reference to the map is passed to functions.
- Maps are commonly used to store data that can be accessed by a unique key.
- Maps are used to implement associative arrays, hash tables, and dictionaries.
- Maps are not thread-safe, meaning that they are not safe for concurrent use.
- Maps are not comparable, meaning that you cannot compare two maps using the `==` operator.
- Maps are reference types, meaning that they are passed by reference to functions.
- Maps are initialized with the zero value of the map type, which is `nil`.