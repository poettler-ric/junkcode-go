package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func main() {
	a := Point{1, 2}
	b := Point{2, 3}

	dictValue := make(map[Point]struct{})
	dictValue[a] = struct{}{}
	dictValue[b] = struct{}{}
	fmt.Println(dictValue)
	if _, ok := dictValue[a]; ok {
		fmt.Println("value: got a by variable")
	}
	if _, ok := dictValue[Point{1, 2}]; ok {
		fmt.Println("value: got a by creation")
	}

	dictReference := make(map[*Point]struct{})
	dictReference[&a] = struct{}{}
	dictReference[&b] = struct{}{}
	fmt.Println(dictReference)
	if _, ok := dictReference[&a]; ok {
		fmt.Println("reference: got a by variable")
	}
	if _, ok := dictReference[&Point{1, 2}]; ok {
		fmt.Println("reference: got a by creation")
	}
}
