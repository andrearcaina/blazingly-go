# Structs

Structs are a way to group related data together. They are similar to classes in other languages, but they are more limited in their functionality. They are useful for grouping data together that is related in some way.

## Defining a Struct

To define a struct, you use the `struct` keyword followed by the name of the struct. Inside the curly braces, you define the fields of the struct. Each field has a name and a type.

```go
package main

type Point struct {
    x, y int
}

type Rectangle struct {
    topLeft, bottomRight Point
}
```

In this example, we define two structs: `Point` and `Rectangle`. The `Point` struct has two fields: `x` and `y`, both of type `int`. The `Rectangle` struct has two fields: `topLeft` and `bottomRight`, both of type `Point`.

## Creating Instances of a Struct

To create an instance of a struct, you use the `var` keyword followed by the name of the variable, the type of the struct, and the values for the fields.

```go
package main

type Point struct {
    x, y int
}

type Rectangle struct {
    topLeft, bottomRight Point
}

func main() {
    var p Point
    p.x = 10
    p.y = 20

    var r Rectangle
    r.topLeft.x = 0
    r.topLeft.y = 0
    r.bottomRight.x = 100
    r.bottomRight.y = 100
}
```

In this example, we create an instance of the `Point` struct called `p` and set the values of its fields. We also create an instance of the `Rectangle` struct called `r` and set the values of its fields.

## Struct Associated Methods

You can define methods on structs in Go. This allows you to define behavior that is associated with a particular struct.

```go
package main

import "fmt"

type Point struct {
    x, y int
}

func (p Point) String() string {
    return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func main() {
    p := Point{10, 20}
    fmt.Println(p.String())
}
```

In this example, we define a method called `String` on the `Point` struct. This method returns a string representation of the `Point` struct. We then create an instance of the `Point` struct called `p` and call the `String` method on it.

Go has a special syntax for defining methods on structs. The method definition starts with the `func` keyword followed by the name of the method, the receiver (in this case, `p Point`), and the return type of the method.

These methods are called associated methods because they are associated with a particular struct. They can access the fields of the struct and modify them if necessary. These are similar to instance methods in other languages. 

However, in Go, methods are not defined inside the struct definition. Instead, they are defined outside the struct definition. This allows you to define methods on structs that are defined in other packages.

Go is also not an object-oriented language, so it does not have classes or inheritance, but similar functionality can be achieved using these associated methods.

## Struct Embedding

You can embed one struct inside another struct in Go. This allows you to reuse the fields and methods of the embedded struct in the outer struct.

```go
package main

import "fmt"

type Point struct {
    x, y int
}

type Circle struct {
    Point
    radius int
}

func main() {
    c := Circle{Point{10, 20}, 5}
    fmt.Println(c.x, c.y, c.radius)
}
```

In this example, we define a `Point` struct and a `Circle` struct that embeds the `Point` struct. This allows us to access the `x` and `y` fields of the `Point` struct directly on the `Circle` struct.

# Interfaces

Interfaces are a way to define behavior in Go. They are similar to interfaces in other languages, but they are more flexible and powerful. They allow you to define a set of methods that a type must implement to satisfy the interface.

## Defining an Interface

To define an interface, you use the `type` keyword followed by the name of the interface and the `interface` keyword. Inside the curly braces, you define the methods that the interface requires.

```go
package main

import "fmt"

type Rectangle struct {
	width, height float64
}

type Triangle struct {
	base, height float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t Triangle) Perimeter() float64 {
	return t.base + 2*t.height
}

func ShapeArea(s Shape) float64 {
	return s.Area()
}

func ShapePerimeter(s Shape) float64 {
	return s.Perimeter()
}

func main() {
	r := Rectangle{10, 20}
	t := Triangle{10, 20}

	fmt.Println(ShapeArea(r))
	fmt.Println(ShapePerimeter(r))

	fmt.Println(ShapeArea(t))
	fmt.Println(ShapePerimeter(t))
}
```

In this example, we define an interface called `Shape` that requires two methods: `Area` and `Perimeter`. We then define two structs: `Rectangle` and `Triangle` that implement the `Shape` interface by defining the `Area` and `Perimeter` methods.

We also define two functions: `ShapeArea` and `ShapePerimeter` that take a `Shape` interface as an argument and call the `Area` and `Perimeter` methods on it.

