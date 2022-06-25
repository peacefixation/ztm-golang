//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a Coordinate // top left
	b Coordinate // bottom right
}

func width(r Rectangle) int {
	return r.b.x - r.a.x
}

func length(r Rectangle) int {
	return r.a.y - r.b.y
}

func area(r Rectangle) int {
	return length(r) * width(r)
}

func perimiter(r Rectangle) int {
	return (width(r) * 2) + (length(r) * 2)
}

func printInfo(r Rectangle) {
	fmt.Println("Area", area(r))
	fmt.Println("Perimiter", perimiter(r))
}

func main() {
	// r := Rectangle{a: Coordinate{1, 4}, b: Coordinate{3, 2}}
	r := Rectangle{a: Coordinate{1, 8}, b: Coordinate{6, 2}}

	printInfo(r)

	r.a.y *= 2
	r.b.x *= 2

	printInfo(r)
}
