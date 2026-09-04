package main

/// Test the behaviour of defer with changing variables.

import (
	"fmt"
)

func main() {
	for i := range 10 {
		defer fmt.Printf("defering: %v\n", i)
	}
}
