package main

import (
	"fmt"
	"gonum.org/v1/gonum/spatial/kdtree"
)

func main() {
	fmt.Println("test")
	points := kdtree.Points{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}
	tree := kdtree.New(points, false)

	queries := kdtree.Points{
		{.7, .4}, // expected result: {1, 0}
		{.4, .4}, // expected result: {0, 0}
		{.6, .6}, // expected result: {1, 1}
		{.4, .9}, // expected result: {0, 1}
		{.1, .1}, // expected result: {0, 0}
		{.1, .9}, // expected result: {0, 1}
		{.9, .1}, // expected result: {1, 0}
		{.9, .9}, // expected result: {1, 1}
	}

	for _, query := range queries {
		found, distance := tree.Nearest(query)
		fmt.Printf("query %v -> %v (%v)\n", query, found, distance)
	}
}
