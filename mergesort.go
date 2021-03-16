package main

import "fmt"

func main() {
	var n = []int{1, 4, 2, 3, 5, 6, 9, 7, 8, 10}
	fmt.Println("unsorted", n)
	var ok = false

	for !ok {
		ok = true
		var i = 0
		for i < len(n)-1 {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
			}
			i++
		}
	}
	fmt.Println("sorted", n)
}
