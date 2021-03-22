package main

import (
	"fmt"
)

func Sort(items []int, first int, last int) {
	if first >= last {
		return
	}
	a := first
	z := last
	x := items[(first+last)/2]

	for a < z {
		for items[a] < x {
			a++
		}
		for items[z] > x {
			z--
		}
		if a <= z {
			var tmp = items[a]
			items[a] = items[z]
			items[z] = tmp
			a++
			z--
		}
	}
	Sort(items, first, z)
	Sort(items, a, last)
}

func quicksort(arr []int) []int {
	length := len(arr)
	items := make([]int, length)
	copy(items, arr)
	Sort(items, 0, length-1)
	return items
}

func main() {
	items := []int{4, 1, 5, 3, 2}

	sortItems := quicksort(items)
	fmt.Println(sortItems)

}
