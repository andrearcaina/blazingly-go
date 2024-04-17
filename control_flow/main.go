package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter any number: ")
	// process input
	var n float64
	_, err := fmt.Scanln(&n)

	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again.")
		fmt.Printf("Error: %v\n", err)
		return
	}

	if n > 0 {
		fmt.Println("The number is positive")
	} else if n < 0 {
		fmt.Println("The number is negative")
	} else {
		fmt.Println("The number is zero")
	}

	// switch statement
	switch n {
	case 0:
		fmt.Println("The number is zero")
	case 1:
		fmt.Println("The number is one")
	case 2:
		fmt.Println("The number is two")
	default:
		fmt.Println("The number is not zero, one or two")
	}

	// switch statement with multiple cases
	switch n {
	case 4, 5, 6:
		fmt.Println("The number is four, five or six")
	default:
		fmt.Println("The number is not four, five or six")
	}

	// switch statement with expression
	switch {
	case n > 0:
		fmt.Println("The number is positive")
		break
	case n < 0:
		fmt.Println("The number is negative")
	default:
		fmt.Println("The number is zero")
	}

	// switch statement with expression and fallthrough
	switch {
	case n > 0:
		fmt.Println("The number is positive")
		fallthrough
	case n < 0:
		fmt.Println("The number is not zero")
	default:
		fmt.Println("The number is zero")
	}

	// what othe control flow statements are there
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	/* for {
		fmt.Println("Infinite loop")
	} */

	// break statement
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}

		fmt.Println(i)
	}

	// continue statement
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}

		fmt.Println(i)
	}

	// goto statement
	goto label
	fmt.Println("This will not be printed")

label:
	fmt.Println("This will be printed")

	// defer statement
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")

	// panic statement
	panic("This is a panic")

	// recover statement (these two statements are used together)
	// these statements are used to handle panics
	// these will not be executed because of the panic statement above
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	panic("This is a panic")

	fmt.Println("This will not be printed")
}
