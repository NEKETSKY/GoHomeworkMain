package main

import (
	"errors"
	"math"
)

//Figure describes 2 methods - area and perimeter of figure
type Figure interface {
	Area() (float64, error)
	Perimeter() (float64, error)
}

//Square describes a square
type Square struct {
	side float64
}

//Circle describes a circle
type Circle struct {
	radius float64
}

var ErrInputSmall = errors.New("bad input. too small")

//area returns the area of a square
func (s Square) Area() (float64, error) {
	if s.side <= 0 {
		return 0, ErrInputSmall
	}
	return s.side * s.side, nil
}

//perimeter returns the perimeter of a square
func (s Square) Perimeter() (float64, error) {
	if s.side <= 0 {
		return 0, ErrInputSmall
	}
	return s.side * 4, nil
}

//area returns the area of a circle
func (c Circle) Area() (float64, error) {
	if c.radius <= 0 {
		return 0, ErrInputSmall
	}
	return math.Pi * c.radius * c.radius, nil
}

//perimeter returns the perimeter of a circle
func (c Circle) Perimeter() (float64, error) {
	if c.radius <= 0 {
		return 0, ErrInputSmall
	}
	return 2 * math.Pi * c.radius, nil
}
