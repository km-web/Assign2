package main

import "fmt"

func main() {
    
    Ages := map[string]int{
        "Kutty": 2,
        "Happy":  1,
        "Scooby":  4,
    }

    dogs := []string{}
    for dog, _ := range Ages {
       dogs = append(dogs, dog)
    }

 
fmt.Println(dogs, Ages)
}