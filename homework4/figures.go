package main

import (
	"fmt"
	"math"
)

//Figure describes 2 methods - area and perimeter of figure
type Figure interface {
	area() float64
	perimeter() float64
}

//Square describes a square
type Square struct {
	side float64
}

//Circle describes a circle
type Circle struct {
	radius float64
}

//area returns the area of a square
func (s Square) area() float64 {
	return s.side * s.side
}

//perimeter returns the perimeter of a square
func (s Square) perimeter() float64 {
	return s.side * 4
}

//area returns the area of a circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//perimeter returns the perimeter of a circle
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	var s Figure = Square{5}
	var c Figure = Circle{5}

	//Case #1. Expected result: area=25, perimeter=20
	if s.area() != 25 {
		fmt.Println("Area of square is incorrect")
	} else if s.perimeter() != 20 {
		fmt.Println("Perimeter of square is incorrect")
	} else {
		fmt.Println("The square is calculated correctly")
	}

	//Case #2. Expected result: area=Pi*25, perimeter=Pi*10
	if c.area() != math.Pi*25 {
		fmt.Println("Area of circle is incorrect")
	} else if c.perimeter() != math.Pi*10 {
		fmt.Println("Perimeter of circle is incorrect")
	} else {
		fmt.Println("The circle is calculated correctly")
	}

	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())

}
