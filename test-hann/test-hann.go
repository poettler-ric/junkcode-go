package main

import (
	"fmt"
	"log"

	"github.com/habedi/hann/core"
	"github.com/habedi/hann/hnsw"
)

func main() {
	fmt.Println("test")
	// Index parameters.
	dim := 2
	m := 5 * 4
	ef := 10 * 4
	distanceName := "euclidean"

	// Create an HNSW index with the given parameters.
	index := hnsw.NewHNSW(dim, m, ef, core.Distances[distanceName], distanceName)
	fmt.Println("Created new HNSW index.")
	fmt.Println(index.Stats())

	// Adding vectors
	vectors := map[int][]float32{
		1: {0, 0},
		2: {0, 1},
		3: {1, 0},
		4: {1, 1},
	}
	for id, vector := range vectors {
		if err := index.Add(id, vector); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Added vectors")
	fmt.Println(index.Stats())

	queries := [][]float32{
		{.7, .4}, // expected result: 3
		{.4, .4}, // expected result: 1
		{.6, .6}, // expected result: 4
		{.4, .9}, // expected result: 2
		{.1, .1}, // expected result: 1
		{.1, .9}, // expected result: 2
		{.9, .1}, // expected result: 3
		{.9, .9}, // expected result: 4
	}
	for _, query := range queries {
		neighbours, err := index.Search(query, 1)
		if err != nil {
			log.Fatal(err)
		}
		for _, n := range neighbours {
			fmt.Printf("query: %v, id: %v, distance: %v\n", query, n.ID, n.Distance)
		}
	}
}
