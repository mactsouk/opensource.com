package main

import (
	"fmt"
)

func change(x [4]int) {
	x[3] = -2
}

func main() {
	anArray := [4]int{-1, 2, 0, -4}
	twoD := [3][3]int{{1, 2, 3}, {6, 7, 8}, {10, 11, 12}}

	fmt.Println("Before change():", anArray)
	change(anArray)
	fmt.Println("After change():", anArray)

	for i := 0; i < len(twoD); i++ {
		k := twoD[i]
		for j := 0; j < len(k); j++ {
			fmt.Print(k[j], " ")
		}
		fmt.Println()
	}

	for _, a := range twoD {
		for _, j := range a {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
