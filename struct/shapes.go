package main

import "math"

type Shape interface {
	Area() float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (circle Circle) Area() float64 {
	return (circle.Radius * circle.Radius) * math.Pi
}

func (triangle Triangle) Area() float64 {
	return triangle.Length * triangle.Height * .5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Triangle struct {
	Length float64
	Height float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}
