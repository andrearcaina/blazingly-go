package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum = 8 // go has type inference
	var bigInt int64 = 1000000
	midInt := 36000 // valid, basically var intNum = 8 without var.

	fmt.Println("------------------------------------")

	fmt.Printf("Regular Integer: %v\n", intNum)
	fmt.Printf("Big Integer: %v\n", bigInt)
	fmt.Printf("Mid Integer: %v\n", midInt)

	intNum += 1

	fmt.Printf("Regular Integer + 1: %v\n", intNum)

	fmt.Println("------------------------------------")

	var floatNum float64 = 12345678.9123
	fmt.Printf("Float Number: %v\n", floatNum)

	fmt.Printf("Big Float Number: %v\n", float64(bigInt)) // casting

	var result float32 = float32(intNum) + float32(midInt)

	fmt.Printf("Result: %v\n", result)

	fmt.Println("------------------------------------")

	var intNum1 int = 3
	var intNum2 int = 2

	fmt.Printf("Integer Division, %v\n", intNum1/intNum2)
	fmt.Printf("Modulo Sign, %v\n", intNum1%intNum2)

	fmt.Println("------------------------------------")

	var myString string = "Hello World!!!!!"
	var backString string = `formatted
weirdly`
	var concatString string = "What" + " " + "Up"

	fmt.Printf("String: %v\n", myString)
	fmt.Printf("Back String: %v\n", backString)
	fmt.Printf("Concat String: %v\n", concatString)
	fmt.Printf("Len of String: %v\n", len(myString))                       // this doesn't actually count the length of the string, but the ascii number of it
	fmt.Printf("Length of String: %v\n", utf8.RuneCountInString(myString)) // this does, but it's weird
	fmt.Printf("Length of 'Rune': %v\n", len("γ"))                         // prints 2
	fmt.Printf("Length of 'Rune': %v\n", utf8.RuneCountInString("γ"))      // prints 1

	// strings are immutable
	// myString[0] = 'h' // this will throw an error

	// strings are actually just a slice of bytes
	// so you can do this
	myString = "hello"
	myString = "c" + myString[1:]

	fmt.Printf("String: %v\n", myString)

	fmt.Println("------------------------------------")

	var Rune rune = 'γ'
	fmt.Printf("Rune: %v\n", Rune) // runes are weird, it's actually just an int32
	fmt.Printf("Len of Rune: %v\n", utf8.RuneLen(Rune))

	fmt.Println("------------------------------------")

	var1, var2 := 1, 2

	fmt.Println(var1, var2)
}
