package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "résumé"

	fmt.Println(s) // prints "résumé"

	for i, c := range s {
		fmt.Println(i, c) // prints 0 114, 1 233, 3 115, 4 117, 5 109, 6 233
		// 114 is the ASCII code for 'r'
		// 233 is the Unicode code for 'é'
		// skips 2 because 'é' is a 2-byte character
		// 105 is the ASCII code for 's'
		// 117 is the ASCII code for 'u'
		// 109 is the ASCII code for 'm'
		// 233 is the Unicode code for 'é'
		// and then adds 1 more location because 'é' is a 2-byte character
	}
	// we can see the above output is not what we expect when printing the length of the string
	// this is because the length of the string is the number of bytes in the string
	fmt.Println(len(s)) // prints 8

	// convert string to slice of runes
	runes := []rune(s)

	fmt.Println(runes) // prints [114 233 115 117 109 233]

	for i, c := range runes {
		fmt.Println(i, c) // prints 0 114, 1 233, 2 115, 3 117, 4 109, 5 233
		// same as above but 's' is not skipped (kinda)
	}

	// convert slice of runes to string
	s = string(runes)

	fmt.Println(s) // prints "résumé"

	var x rune = 'a'
	var y rune = 'b'

	fmt.Println(x, y)           // prints 97 98
	fmt.Printf("%T %T\n", x, y) // prints int32 int32

	var z byte = 'c'

	fmt.Println(z) // prints 99

	// string building
	var strSlice = []string{"h", "e", "l", "l", "o", ",", " ", "w", "o", "r", "l", "d"}
	var str string

	for i := range strSlice {
		str += strSlice[i]
	}

	fmt.Println(str) // prints "hello, world" (very inefficient as it creates a new string every time we add a character)

	// better way to build strings

	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	fmt.Println(strBuilder.String()) // prints "hello, world" (much more efficient as it doesn't create a new string every time we add a character)

	// slicing a string

	newString := "hello, world"
	s2 := newString[0:5] // "hello"

	fmt.Println(newString) // "hello, world"
	fmt.Println(s2)        // "hello"

	s3 := newString[0]

	fmt.Println(s3) // 104 (ASCII code for 'h')

	// to convert s3 to a string we can do this
	s4 := string(s3)

	fmt.Println(s4) // "h" (this creates a new string that is a copy of the byte at index 0 of newString)
}
