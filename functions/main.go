package main

import (
	"fmt"
)

func printAnything(arg string) {
	fmt.Println(arg)
}

func sumN(n int) int {
	sum := 0

	for i := 0; i < n; i++ {
		sum += i
	}

	return sum
}

func sumNRecursive(n int) int {
	if n == 0 {
		return 0
	}

	return n + sumNRecursive(n-1)
}

func doSomething(n, m int) (int, int) {
	return n + m, n * m
}

func main() {
	fmt.Println("------------------------------------")

	printAnything("Hello World")

	fmt.Println("------------------------------------")

	fmt.Printf("Sum of %v = %v\n", 10, sumN(10))
	fmt.Printf("Sum of %v = %v\n", 10, sumNRecursive(10))

	fmt.Println("------------------------------------")

	x, y := doSomething(10, 20)

	fmt.Printf("Sum: %v, Product: %v\n", x, y)

	fmt.Println("------------------------------------")

	canAlsoGoHere()

	fmt.Println("------------------------------------")

	func() {
		fmt.Println("I am an anonymous function")
	}()

	func(x int) {
		fmt.Printf("I am an anonymous function with argument %v\n", x)
	}(10)

	z := func(x int) int {
		return x * x
	}

	fmt.Printf("I am an anonymous function with return value %v\n", z(10))

	fmt.Println("------------------------------------")
}

func canAlsoGoHere() {
	fmt.Println("I can also go here")
}
