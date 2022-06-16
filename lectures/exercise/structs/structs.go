//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:


package main

import "fmt"

//* Create a rectangle structure containing its coordinates
type Rectangle struct {
	Width int
	Length int
}

//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
func calculateArea(coordinates Rectangle) int {
	return coordinates.Length * coordinates.Width
}

func calculatePerimeter(coordinates Rectangle) int{
	return (coordinates.Length * 2) + (coordinates.Width * 2)
}


//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal


//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

func main() {
	var sampleRectangle = Rectangle{3,5}
	fmt.Println(calculateArea(sampleRectangle))
	fmt.Println(calculatePerimeter(sampleRectangle))

	sampleRectangle.Length *= 2
	sampleRectangle.Width *= 2

	fmt.Println(calculateArea(sampleRectangle))
	fmt.Println(calculatePerimeter(sampleRectangle))

}
