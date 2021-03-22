package main

import (
	"fmt"
)

func main() {

	items := []int{4, 1, 5, 3, 2}

	fmt.Println("Unsorted", items)
	fmt.Println("Sorted", mergeSort(items))
}

func merge(left, right []int) (result []int) {

	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

func mergeSort(items []int) []int {
	var length = len(items)
	middle := int(length / 2)
	var left = make([]int, middle)
	var right = make([]int, length-middle)

	if length == 1 {
		return items
	}

	for i := 0; i < length; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}
