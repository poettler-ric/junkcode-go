package main

import (
	"fmt"
)

func main() {
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers)
	sub := numbers[1:4]
	fmt.Println(sub)
	sub[2] = 11
	fmt.Println(sub)
	fmt.Println(numbers)
}
