package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) GetDistanceFrom(point *Point) float64 {
	return math.Sqrt(math.Pow(p.x-point.x, 2) + math.Pow(p.y-point.y, 2))
}

func main() {
	var x float64
	var y float64

    fmt.Println("Enter x and y for point 1")
	if _, err := fmt.Scanln(&x, &y); err != nil {
		fmt.Println("Incorrect input format")
		return
	}

	p1 := NewPoint(x, y)

    fmt.Println("Enter x and y for point 2")
	if _, err := fmt.Scanln(&x, &y); err != nil {
		fmt.Println("Incorrect input format")
	}

	p2 := NewPoint(x, y)

    fmt.Println("Distance:", p1.GetDistanceFrom(p2))
}
