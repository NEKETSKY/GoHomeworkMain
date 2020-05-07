package main

import "fmt"

//type contains 2 variable point coordinates
type Point struct {
	x, y int
}

//type contains the coordinate of the starting point of the square and the length of its side
type Square struct {
	start Point
	a     uint
	//"start" variable contains the starting point of type point, "a" variable contains the length of the side of the square
}

//returns the end point of the square
func (s Square) End() Point {
	return Point{s.start.x + int(s.a), s.start.y - int(s.a)}
}

//returns the perimeter of a square
func (s Square) Perimeter() uint {
	return s.a * 4
}

//returns the square area
func (s Square) Area() uint {
	return s.a * s.a
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
