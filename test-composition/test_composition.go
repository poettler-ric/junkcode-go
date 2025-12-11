package main

import (
	"fmt"
)

type Point struct {
	X, Y float64
}

func (p Point) String() string {
	return fmt.Sprintf("Point{%v %v}", p.X, p.Y)
}

type WeightedPoint struct {
	*Point
	Weight float64
}

func (p WeightedPoint) String() string {
	return fmt.Sprintf("WeightedPoint{%v %v, weight: %v}", p.X, p.Y, p.Weight)
}

func main() {
	p := Point{5, 6}
	wp := WeightedPoint{&Point{2, 3}, 9}

	fmt.Println(p)
	fmt.Println(wp)
	// direct access to wp.Point.X
	fmt.Printf("wp.X: %v <-> wp.Point.X: %v\n", wp.X, wp.Point.X)

}
