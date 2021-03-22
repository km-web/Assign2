package main

import "fmt"

func Permutation(numRune []rune, left, right int) {
	j := numRune[left]

	if left == right {
		fmt.Println(string(numRune))
	} else {
		for k := left; k <= right; k++ {
			j, numRune[k] = numRune[k], j
			Permutation(numRune, left+1, right)
			j, numRune[k] = numRune[k], j
		}
	}
}
func main() {
	num := "1234"
	numRune := []rune(num)
	Permutation(numRune, 0, len(numRune)-1)
}
