package main

import (
	"fmt"
)

func main() {

	array := []int{1, 2, 3, 4, 5, 6}
	num := 2
	for i := 0; i < 6; i++ {
		if array[i] == num {
			fmt.Println("Position : ", i+1)
			return

		}

	}
}
