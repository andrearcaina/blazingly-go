package main

import (
	"custom_types/src"
	"fmt"
)

func main() {
	fmt.Println("--------------------------------------------------")

	var myEngine = src.New(20, 10)

	myEngine.SetOwnerInfo("John Doe", 25)

	fmt.Printf("Type of myEngine: %T\n", myEngine)
	fmt.Printf("MPG: %v\n", myEngine.GetMpg())
	fmt.Printf("Gallons: %v\n", myEngine.GetGallons())
	fmt.Printf("Owner Info: %v\n", myEngine.GetOwnerInfo())
	fmt.Printf("Owner Name: %v\n", myEngine.GetOwnerName())
	fmt.Printf("Owner Age: %v\n", myEngine.GetOwnerAge())
	fmt.Printf("Miles Left: %v\n", myEngine.MilesLeft())

	fmt.Println("--------------------------------------------------")

	myEngine.SetGallons(30)

	fmt.Println("I'm refueling!")

	fmt.Printf("MPG: %v\n", myEngine.GetMpg())
	fmt.Printf("Gallons: %v\n", myEngine.GetGallons())

	fmt.Println("I'm driving!")

	myEngine.Drive()

	fmt.Printf("Miles Left: %v\n", myEngine.MilesLeft())
	fmt.Printf("Memory Location of myEngine: %v\n", &myEngine)

	if src.CanMakeIt(myEngine) {
		fmt.Println("I can make it!")
	} else {
		fmt.Println("I can't make it!")
	}

	fmt.Println("--------------------------------------------------")

	// not recommended to use anonymous structs
	// this is just to show how to create one
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}

	fmt.Printf("Type of myEngine2: %T\n", myEngine2)
	fmt.Printf("MPG: %v\n", myEngine2.mpg)
	fmt.Printf("Gallons: %v\n", myEngine2.gallons)

	fmt.Println("--------------------------------------------------")

	fmt.Println("Using Interfaces")

	r := src.Rectangle{Width: 10, Height: 20}
	t := src.Triangle{Base: 10, Height: 20}

	fmt.Printf("Area of Rectangle: %v\n", src.ShapeArea(r))
	fmt.Printf("Perimeter of Rectangle: %v\n", src.ShapePerimeter(r))

	fmt.Printf("Area of Triangle: %v\n", src.ShapeArea(t))
	fmt.Printf("Perimeter of Triangle: %v\n", src.ShapePerimeter(t))
}
