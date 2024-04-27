package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 3}
	floatSlice := []float32{1, 2, 3}

	fmt.Println(isEmpty(intSlice))
	fmt.Println(isEmpty(floatSlice))
	fmt.Println(sumSlice[int](intSlice))
	fmt.Println(sumSlice[float32](floatSlice))
}

// sumSlice sums the elements of a slice of integers or floats
// T is a "generic" type that can be either int, float32, or float64 using the type constraint syntax (|)
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T = 0

	for _, v := range slice {
		sum += v
	}

	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
