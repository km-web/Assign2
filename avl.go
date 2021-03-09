package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	start := 0
	end := len(a) - 1
	mid := (start + end) / 2
	x := 6

	if mid < len(a) || a[mid] == x {
		fmt.Printf("found %d at index  %v\n", x, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}

}
