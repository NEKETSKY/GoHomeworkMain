package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
type Shape interface {
	Area() (float64, error)
	Perimeter() (float64, error)
}

func TestArea(t *testing.T) {
	tests := map[string]struct {
		shape Shape
		out   float64
		err   error
	}{
		"square 10": {Square{10}, 100, nil},
		"square 0":  {Square{0}, 0, ErrInputSmall},
		"circle 10": {Circle{10}, 314.1592653589793, nil},
		"circle 0":  {Circle{0}, 0, ErrInputSmall},
	}

	for name, tt := range tests {
		res, err := tt.shape.Area()
		assert.Equal(t, res, tt.out, "%s: result - %v; expected result - %v", name, res, tt.out)
		assert.Equal(t, err, tt.err, "%s: errors must be equal", name)
	}
}

func TestPerimeter(t *testing.T) {
	tests := map[string]struct {
		shape Shape
		out   float64
		err   error
	}{
		"square 10": {Square{10}, 40, nil},
		"square 0":  {Square{0}, 0, ErrInputSmall},
		"circle 10": {Circle{10}, 62.83185307179586, nil},
		"circle 0":  {Circle{0}, 0, ErrInputSmall},
	}

	for name, tt := range tests {
		res, err := tt.shape.Perimeter()
		assert.Equal(t, res, tt.out, "%s: result - %v; expected result - %v", name, res, tt.out)
		assert.Equal(t, err, tt.err, "%s: errors must be equal", name)
	}
}
