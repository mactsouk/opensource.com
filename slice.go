package main

import (
	"fmt"
)

func negative(x []int) {
	for i, k := range x {
		x[i] = -k
	}
}

func printSlice(x []int) {
	for _, number := range x {
		fmt.Printf("%d ", number)
	}
	fmt.Println()
}

func main() {
	s := []int{0, 14, 5, 0, 7, 19}
	printSlice(s)
	negative(s)
	printSlice(s)

	fmt.Printf("Before. Cap: %d, length: %d\n", cap(s), len(s))
	s = append(s, -100)
	fmt.Printf("After. Cap: %d, length: %d\n", cap(s), len(s))
	printSlice(s)

	anotherSlice := make([]int, 4)
	fmt.Printf("A new slice with 4 elements: ")
	printSlice(anotherSlice)
}
