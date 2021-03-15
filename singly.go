package main

import (
	"container/list"
	"fmt"
)

func main() {

	values := list.New()
	values.PushFront(10)

	for i := 0; i < 5; i++ {

		values.PushBack(i)

	}

	for List := values.Front(); List != nil; List = List.Next() {
		fmt.Println(List.Value)
	}
	values.PushBack(20)
	for List := values.Back(); List != nil; List = List.Next() {
		fmt.Println(List.Value)
	}
}
