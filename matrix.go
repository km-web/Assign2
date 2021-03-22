package main

import "fmt"

func main() {
	var i int
	var j int
	var k int

	for i = 1; i < 2; i++ {
		for j = 1; j < 2; j++ {
			for k = 1; k < 4; k++ {

				fmt.Println(i, j, k)

			}
		}

	}
}
