# Strings, Runes, Bytes (UTF-8 encoding)

## Strings
- A string is a read-only slice of bytes.
- Strings are immutable.
- A string holds arbitrary data.
- A string is a collection of runes.
- A string is a collection of bytes.
- A string is a collection of characters.

Golang strings are made up of bytes, but those bytes do not define the character encoding of the string. The character encoding is defined by the Unicode standard. The Go language uses UTF-8 encoding for strings. UTF-8 is a variable-length encoding that uses 1 to 4 bytes to represent each Unicode code point.

An example of a string in Go is "Hello, 世界". This string contains 13 bytes, but only 9 characters. The first 7 bytes represent the ASCII characters "Hello, ", and the last 6 bytes represent the two Unicode characters "世" and "界".

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
    s := "Hello, 世界"
    fmt.Println(s)
    fmt.Println(len(s))
    fmt.Println(utf8.RuneCountInString(s))
}

// Output:
// Hello, 世界
// 13 (representing 13 bytes in the string)
// 9 (representing 9 characters in the string)
```

## Runes
- A rune is a Unicode code point.
- A rune is an alias for the `int32` type.
- A rune is a single Unicode character.

```go
package main

import (
    "fmt"
)

func main() {
    r := '世'
    fmt.Println(r)
}

// Output:
// 19990
```

## Bytes
- A byte is an alias for the `uint8` type.

```go
package main

import (
    "fmt"
)

func main() {
    a := byte(0x41)
    b := 'A'
    fmt.Println(a)
    fmt.Println(b)
}

// Output:
// 65
// 65
```
