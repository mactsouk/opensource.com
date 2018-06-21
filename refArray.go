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
	anArray := [5]int{-1, 2, -3, 4, -5}
	refAnArray := anArray[:]

	fmt.Println("Array:", anArray)
	printSlice(refAnArray)
	negative(refAnArray)
	fmt.Println("Array:", anArray)
}
